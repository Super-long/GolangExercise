package main

import (
	"io/ioutil"
	"net/http"
)

// 一个支持并发的缓存(函数记忆)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}	// 通知数据何时获取完成
}

type request struct {
	key      string
	response chan<- result // 用于传输结果
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func){
	cache := make(map[string]*entry)
	for req := range memo.requests{
		e := cache[req.key]
		if e == nil{
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// f函数返回时退出阻塞
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

func httpGFetBody(url string)(interface{}, error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main(){
	//m := New(httpGFetBody)
	// 后面直接m.Get就可以了

}