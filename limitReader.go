// io包中的LimitReader函数接受一个io.Reader r和字节数n，返回一个Reader，
// 该返回值从r读取数据，但在读取n字节后报告文件结束，实现该函数。
package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	reader  io.Reader
	limit   int
	current int
}

func (r *LimitedReader) Read(b []byte) (int, error) {
	if r.current >= r.limit {
		return 0, io.EOF
	}
	var flag bool
	if r.current+len(b) > r.limit {
		b = b[:r.limit-r.current]
		flag = true
	}
	n, err := r.reader.Read(b)
	if err != nil {
		return n, err
	}
	r.current += n
	if flag {	// 当读取大于limit时直接报EOF
		err = io.EOF
	}
	return n, err
}

func LimitReader(r io.Reader, li int) io.Reader {
	// 可能传入的li小于r中的数据量
	res := LimitedReader{
		reader: r,
		limit:  li,
	}
	return &res
}

func main(){
	file, err := os.Open("README.md")
	if err != nil{
		panic(err)
	}
	lr := LimitReader(file, 12)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if closeErrr := file.Close(); err != nil{ // go程序设计语言P115
		err = closeErrr
		panic(err)
	}
	fmt.Println(n, buf)
}