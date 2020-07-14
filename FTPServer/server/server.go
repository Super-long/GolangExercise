package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"exercise/FTPServer/server/ftp"
	"exercise/FTPServer/ftp"
)

func handleFunc(con net.Conn) {
	defer con.Close()

	// 验证的协议为长度+内容
	// 身份验证
	// 读取用户名
	var length uint32
	// 先读取用户名长度
	err := binary.Read(con, binary.LittleEndian, &length)
	if err != nil {
		err = binary.Write(con, binary.LittleEndian, uint32(0))
		if err != nil {
			log.Println(err)
		}
		return
	}
	// 用户发来的实际数据
	user := make([]byte, length-uint32(binary.Size(length)))
	err = binary.Read(con, binary.LittleEndian, user)
	if err != nil {
		err = binary.Write(con, binary.LittleEndian, uint32(0))
		if err != nil {
			log.Println(err)
		}
		return
	}

	// 读取密码
	err = binary.Read(con, binary.LittleEndian, &length)
	if err != nil {
		err = binary.Write(con, binary.LittleEndian, uint32(0))
		if err != nil {
			log.Println(err)
		}
		return
	}
	pwd := make([]byte, length-uint32(binary.Size(length)))
	err = binary.Read(con, binary.LittleEndian, pwd)
	if err != nil {
		err = binary.Write(con, binary.LittleEndian, uint32(0))
		if err != nil {
			log.Println(err)
		}
		return
	}

	// 验证用户名密码获取家目录
	validated, cwd := server.Validate(ftp.Sbyte2str(user), ftp.Sbyte2str(pwd))
	if !validated {
		err = binary.Write(con, binary.LittleEndian, uint32(0))
		if err != nil {
			log.Println(err)
		}
		return
	}

	// cwd为该用户的家目录,将家目录返回给客户端,先写长度再写数据
	home := ftp.Str2sbyte(cwd)
	err = binary.Write(con, binary.LittleEndian, uint32(binary.Size(home)))
	if err != nil {
		log.Println(err)
		return
	}
	err = binary.Write(con, binary.LittleEndian, home)
	if err != nil {
		log.Println(err)
		return
	}

	ftpCon := ftp.FtpConn{
		Con:  con,
		Home: cwd,
		Cwd:  cwd,
	}
	ftpServer := server.FtpServer{
		ftpCon,
	}
	// 循环监听命令请求
	for !ftpServer.Exit {
		var length uint32
		err = binary.Read(con, binary.LittleEndian, &length)
		if err != nil {
			log.Println(err)
			return
		}
		// 读取命令的编号
		var cmdid uint8
		fmt.Printf("当前接收到的命令为 %d\n", cmdid)
		err = binary.Read(con, binary.LittleEndian, &cmdid)
		if err != nil {
			log.Println(err)
			return
		}
		// 读取命令的参数
		args := make([]byte, length-uint32(binary.Size(cmdid))-uint32(binary.Size(length)))
		err = binary.Read(con, binary.LittleEndian, args)
		if err != nil {
			log.Println(err)
			return
		}

		switch cmdid {
		case ftp.Commands["cd"]:
			err = ftpServer.HandleCd(args)
		case ftp.Commands["ls"]:
			err = ftpServer.HandleLs(args)
		case ftp.Commands["exit"]:
			err = ftpServer.HandleExit(args)
		case ftp.Commands["mkdir"]:
			err = ftpServer.HandleMkdir(args)
		case ftp.Commands["put"]:
			err = ftpServer.HandlePut(args)
		case ftp.Commands["get"]:
			err = ftpServer.HandleGet(args)
		default:
			err = ftpServer.Write([]byte("no command handler."))
		}

		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	server.Init() // 初始化用户信息,搞成数据库也可,晚上再说
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleFunc(con) // 可能会导致同时执行的goroutine太多,可以搞一个令牌池限定下同时处理的数量
	}
}