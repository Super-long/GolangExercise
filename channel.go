package main

import (
	"fmt"
	"time"
)

/*
1.channel的读写操作都是阻塞的
2.channel是同步的,意味着同一时间最多只能有一个goroutine来操作
 */

func main() {
	/*
	var a chan string
	fmt.Printf("%T,%v\n",a,a)

	if a == nil{
		fmt.Println("can`t use.")
		a = make(chan string)
		fmt.Println(a)
	}
	test1(a) //输出与上面相同，证明是一个引用类型的传递
	*/

	/*
	var b chan bool
	b = make(chan bool)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println("i:",i)
		}
		b <- true //向通道中发送数据 <-其实就是数据的流向
		fmt.Println("结束")
	}()
	data := <-b //data从channel中读取数据 这样看起来比较类似于std::promise或者条件变量,但更加强大
	fmt.Println(data)
	fmt.Println("main -- over")
	*/

	testBlocking()

	ch := make(chan int)
	ch <- 10 //这样就会出现死锁 无论是读了没人写的channel还是写了没人读的channel都会发生死锁
}

func test1(ch chan string){
	fmt.Printf("%T,%v\n",ch,ch)
}

func testBlocking(){
	ch1 := make(chan int)
	go func(){
		fmt.Println("子goroutine开始执行..")
		//time.Sleep(3*time.Second)
		data := <-ch1
		fmt.Println("data :",data)
	}()

	time.Sleep(3*time.Second)
	ch1 <- 10
	time.Sleep(1*time.Second) //使子goroutine能够读到数据
	fmt.Println("main .. over ")
}