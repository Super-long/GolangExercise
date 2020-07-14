package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

var (
	message = make(chan string, 10)
	enter   = make(chan client)
	leave   = make(chan client)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-enter:
			clients[cli] = true
		case cli := <-leave:
			delete(clients, cli)
			close(cli)
		case msg := <-message:
			for cli := range clients { // 广播信息
				// 一步很慢的话可以导致所有的用户卡住 所以要注意handleConnect中的ch，应改为缓冲且写入非阻塞
				// 书上意思是当非阻塞写入失败的时候直接跳过，但是其实应该搞一个缓冲区，存放未发出的消息，过一阵再发
				// 可以使用select中的default来实现非阻塞读写
				select {
				case cli <- msg:
				default:	// 指定没有其他的通信发生的时候可以立即执行的动作
					log.Println("阻塞， 跳过此信息")
				}
			}
		}
	}
}

// 8.13 需要断掉空置连接
// 在另一个goroutine中调用conn.Close可以让当前阻塞的读操作变得非阻塞
func handleConnect(conn net.Conn) {
	ch := make(chan string,10)
	done := make(chan struct{})

	go clientWriter(conn, ch)

	go func() {	// 用作关闭空闲连接
		for {
			select {
			case <-time.After(7 * time.Second):
				conn.Close()
				fmt.Println("超时了")
				return
			case <-done:
				// 什么也不用做
			}
		}
	}()

	who := conn.RemoteAddr().String()
	ch <- "You are " + who          // 显示给自己看的
	message <- who + "has arrived." // 广播的
	enter <- ch

	input := bufio.NewScanner(conn)
	// 客户端断开是从for循环跳出
	for input.Scan(){ // 客户端发送给这个套接字的信息
		fmt.Println("nihao")
		message <- who + " : " + input.Text()
		done <- struct{}{}
	}

	leave <- ch
	message <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch { // 此用户接收到的信息 需要发送给客户端
		// 使msg格式化后写入conn 这一步可能很慢，导致客户无法写入
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnect(conn)
	}
}
