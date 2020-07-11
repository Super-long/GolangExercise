package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
func (d dollars) String() string {return fmt.Sprintf("$%.2f", d)}

type database map[string]dollars
/*func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
	fmt.Println("up\n")
	for item, price := range db {
		//fmt.Println(item, price)
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
	//fmt.Println("down\n")
}*/

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s : %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

/*func main(){
	db := database{"shoes" : 50, "socks" : 5}
	// database已经实现了http.Request接口
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}*/

func main(){
	// 请求多工转发器ServeMux,用来简化URL和处理程序之间的关联,
	// 如果不使用的话只能使用使用ServeHTTP作为处理时的回调了
	// 这样的话ServeHTTP就过于臃肿了
	db := database{"shoes" : 50, "socks" : 5}
	mux := http.NewServeMux()	// 实现了handler的接口
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}