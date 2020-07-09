package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {	// 延迟关闭打开的文件
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main(){
	name, len, err := fetch("https://golang.org")
	if err != nil{
		log.Print(err)
		return
	}
	fmt.Println(name, len)
}