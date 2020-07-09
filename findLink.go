package main

import (
	"fmt"
	"log"
	"net/http"
	//"github.com/golang/net/html"
	"golang.org/x/net/html"
	"os"
)

// 类似于一个爬虫，通过一个网页，爬取多个网站

// 可以根据传入的url解析出更多的url
func Extract(url string) ([]string, error) {
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
	fmt.Println("up")
	forEachNode(doc, visitNode, nil)
	return links, nil
}


func forEachNode(n *html.Node, pre, post func(n *html.Node)){
	if pre != nil{
		pre(n)
	}

	for c:=n.FirstChild; c!= nil; c = c.NextSibling{
		fmt.Println("lop")
		forEachNode(c, pre, post)
	}
	if post != nil{
		post(n)
	}
}

// 爬虫主体 对传入的url调用Extract， 得到其返回的信息中的所有url
func crawl(url string) []string{
	fmt.Println(url)
	list, err := Extract(url)
	fmt.Println("函数返回")
	if err != nil{
		log.Print(err)
	}
	return list
}

// 对worklist中每一个元素调用f，并把结果集加入worklist
func breadthFirst(f func(item string) []string, worklist []string){
	seen := make(map[string]bool)
	for len(worklist) > 0{
		items := worklist
		worklist = nil
		for _, item := range items{
			if !seen[item]{
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main(){
	breadthFirst(crawl, os.Args[1:])
}