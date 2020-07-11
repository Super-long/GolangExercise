package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func fetch_(url string) (filename string, n int64, err error) {
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
/*	name, len, err := fetch_("http://127.0.0.1:8000")
	if err != nil{
		log.Print(err)
		return
	}
	fmt.Println(name, len)*/
	// 与fetchserver配套使用
	// "http://127.0.0.1:8000/list"
	// "http://127.0.0.1:8000/price?item=socks"
	resp ,err := http.Get("http://127.0.0.1:8000/price?item=hat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b,err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch reading  %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}