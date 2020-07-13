package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {	// 创建abort通道
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	trick := time.Tick(1 * time.Second) // 相当于一个定时器
	fmt.Println("countdown!")
	// 如果多个情况同时满足的话，seclect随机选择一个，这样会使得程序运行难以确定
	for countdown := 10; countdown>0; countdown--{
		select {
		case <- trick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	fmt.Println("launch sucessful!")
}
