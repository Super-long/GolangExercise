package main

import (
	"fmt"
	"sync"
	"time"
)

/*
返回值被忽略
*/
var wg sync.WaitGroup
var rwMutex *sync.RWMutex

func main() {
	/*go printNum()
	for i:=1;i<=1000;i++{
		fmt.Println("\tgoroutine中打印字母: A\n",i)
	}
	time.Sleep(1*time.Second) //可能主函数执行完了 goroutine还没有完毕
	*/

	//------------------

	/*a := 1
	go func (){
		a = 2
		fmt.Println("goroutine")
	}()
	a = 3
	time.Sleep(1*time.Second)
	fmt.Println(a)*/

	// ---------------------
	// waitGroup 同步等待组
	/*wg.Add(2)
	go printNum()
	go printletter() //不会执行完 因为主goroutine先执行完

	fmt.Println("等待wg中的子goroutine执行完才可以从阻塞中出去")
	wg.Wait()*/

	// ---------------------
	rwMutex = new(sync.RWMutex)

	//wg.Add(2)
	//go readDate(1)
	//go readDate(2)
	//wg.Wait()

	wg.Add(3)
	go writeData(1)
	go readDate(1)
	go writeData(2)
	fmt.Println("main over")
	wg.Wait()
}

func readDate(i int) {
	defer wg.Done()
	fmt.Println("开始读:")
	rwMutex.RLock()
	fmt.Println(i, "正在读取数据")
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	rwMutex.RLocker()
	fmt.Println("读结束")
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println("开始写操作")
	rwMutex.Lock()
	fmt.Println("正在写")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
}

func printNum() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("goroutine中打印数字: %d\n", i)
	}
	wg.Done() //计数器减一
}

func printletter() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		fmt.Println("goroutine中打印字母: A\n")
	}
}
