package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn){
	defer c.Close()
	for{
		// http://docscn.studygolang.com/pkg/time/#Time.Format 关于format的参数
		_,err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil{
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		// 不启动goroutine的话第二个连接会阻塞，因为服务器没有监听套接字
		go handleConn(conn)
	}
}