package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		子goroutine每写一次,主gouroutine就会读取一次 解除阻塞

		主goroutine,读取数据
			每次读取数据前都会阻塞,子goroutine写入后解除阻塞
	*/
	ch1 := make(chan int)
	go send(ch1)
	/*for {
		time.Sleep(1*time.Second)
		v, ok := <-ch1
		if !ok { //最后一次读取的时候 v无意义 一般是对应类型的默认值
			fmt.Println("已经读取所有的数据")
			break
		}
		fmt.Println("读取的数据是:",v)
	}*/
	for v := range ch1 { //相比于上面更为简单的写法
		fmt.Println("读取到的数据是", v)
	}

	//-----------缓冲通道
	/*
		非缓冲通道 make(chan T)
			一次发送 一次接收 都是阻塞的
		缓冲通道   make(chan T, capacity)
			发送仅当缓冲区满了 才会阻塞
			接收只有当缓冲区为空了才会阻塞
	*/
	ch2 := make(chan int)
	fmt.Println(len(ch2), cap(ch2)) //0 0

	ch3 := make(chan int, 5)
	fmt.Println(len(ch3), cap(ch3)) //0 5
	ch3 <- 100
	fmt.Println(len(ch3), cap(ch3)) //1 5

	// -------------------------------------------------单向通道
	ch4 := make(chan int)   //双向
	ch5 := make(chan<- int) //单向 只能写不能读
	//ch6 := make(<-chan int) // 单向 只能读不能写

	//ch4 <- 100
	//data4 := <-ch4 //双向通道 所以都可以

	//ch5 <- 1000
	//data4 := <-ch5 //invalid operation: <-ch5 (receive from send-only type chan<- int)

	//ch6 <- 100 //invalid operation: ch6 <- 100 (send to receive-only type <-chan int)
	//data6 := <-ch6

	go fun1(ch4)
	go fun1(ch5)
	// 单向通道直接创建出来感觉意义不是特别大 用处就是我们可以创建一个双向的通道 然后去匹配一个单向管道的参数 限制这个函数的作用
	// 比如fun1 就只能写而不能读

	d4 := <-ch4
	fmt.Println("data is ",d4)
	time.Sleep(1*time.Second)
}

func send(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1) //去掉的话在主goroutine中出现死锁,读取的channel没有写入了
}

//该函数只能操作只写的通道
func fun1(ch chan <- int){
	ch <- 100
	fmt.Println("写入完毕")
}