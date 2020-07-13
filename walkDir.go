package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 并发遍历目录

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} // 默认遍历当前路径
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, entry := range roots {
		n.Add(1)
		go walkDir(entry, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes) // 遍历结束以后关闭filesizes
	}()

	var nfiles, nbytes int64 // 记录文件总字节数

	trick := time.Tick(500 * time.Millisecond)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)	// 从标准输入接收数据
	}()
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-trick: // 逐秒输出，显示友好一点
			fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
		}
	}
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9) // 不加可能最后会漏一次输出
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// 遍历目录
func walkDir(dir string, n *sync.WaitGroup, filesizes chan<- int64) {
	defer n.Done()
	if cancelled(){
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, filesizes)
		} else {
			filesizes <- entry.Size()
		}
	}
}

// 限制同时最多运行20个goroutine
var sema = make(chan struct{}, 20)

// 返回dir目录下的所有文件信息
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //得到令牌
	defer func() { <-sema }() // TODO 为什么这么写

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dirents : %v\n", err)
		return nil
	}
	return entries
}
