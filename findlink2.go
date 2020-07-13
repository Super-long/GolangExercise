package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"os"
)

// 相比于findlink，这个是一个并发的爬虫

func Extract_(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("出现错误")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode_(doc, visitNode, nil)
	return links, nil
}

func forEachNode_(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode_(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// 令牌，发送一个值领取令牌。取出一个值放弃令牌，用来限制同时执行的Extract_
var token = make(chan struct{}, 20)

// 爬虫主体 对传入的url调用Extract， 得到其返回的信息中的所有url
/*func crawl_(url string) []string{ // 版本1
	fmt.Println(url)
	list, err := Extract_(url)
	if err != nil{
		log.Print(err)
	}
	return list
}*/

// 使用令牌限制通知执行的数量
func crawl_(url string) []string { // 版本1
	fmt.Println(url)
	token <- struct{}{} // 获取令牌
	list, err := Extract_(url)
	<-token // 释放令牌
	if err != nil {
		log.Print(err)
	}
	return list
}

// main函数相当于findlink中的breadthFirst
// 版本1
// 1.无法在搜索完毕后正常退出，
// 可能使得同时运行的goroutine太多
/*func main(){
	wordlist := make(chan []string)
	go func(){
		wordlist <- os.Args[1:]
	}()
	seen := make(map[string]bool) // 相当于一个set，去重
	for list := range wordlist{
		for _,link := range list{
			if !seen[link]{
				seen[link] = true
				go func(link string){
					wordlist <- crawl_(link)
				}(link)
			}
		}
	}
}*/

func main() {
	wordlist := make(chan []string)
	var n int
	n++
	go func() {
		wordlist <- os.Args[1:]
	}()
	seen := make(map[string]bool) // 相当于一个set，去重
	for ; n > 0; n-- {	// n的值实际上是wordlist中的值
		list := <-wordlist
		for _,link := range list{
			if !seen[link]{
				seen[link] = true
				n++
				go func(link string){
					wordlist <- crawl_(link)
				}(link)
			}
		}
	}
}
