package ftp

import (
	"encoding/binary"
	"net"
	"unsafe"
)

var Commands = map[string]uint8{
	"cd":    uint8(1),
	"ls":    uint8(2),
	"exit":  uint8(3),
	"mkdir": uint8(4),
	"put":   uint8(5),
	"get":   uint8(6),
}

type FtpConn struct {
	Con  net.Conn	// 维护的连接实体
	Cwd  string		// 当前路径
	Home string		// 用户目前的家目录
	Exit bool		// 是否退出
}

func (ftpCon *FtpConn) Write(content []byte) error {
	var length uint32
	length = uint32(len(content))
	if length == 0 {
		return binary.Write(ftpCon.Con, binary.LittleEndian, &length)
	}
	length = length + uint32(binary.Size(length))	// 写入全部的数据后再写入一个内容的长度
	err := binary.Write(ftpCon.Con, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	err = binary.Write(ftpCon.Con, binary.LittleEndian, content)
	if err != nil {
		return err
	}
	return nil
}

// string转[]byte
// 利用string本来的底层数组
// go圣经 p268
func Str2sbyte(s string) (b []byte) {
	/*
	string 的底层数组结构如下
	struct string {
	unit8 *str
	int len
	}

	而 []byte 的底层结构如下
	struct uint8 {
	unit8 *array
	int len
	int cap
	}
	*/
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

// []byte转string
// 利用[]byte本来的底层数组
func Sbyte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}