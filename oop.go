package main

import "fmt"

func main() {
	p1 := Person{"zhangsan",16}
	fmt.Println(p1)
	s1 := Student{p1,"weibing"}
	fmt.Println(s1)

	s2 := Student{Person:Person{"zhaolong",16},school: "beijindaxue"}
	fmt.Println(s2)

	s2.name = "nihao" //两种方法均可
	s2.Person.name = "hello"
	fmt.Println(s2)
	// -----------方法
	w1 := worker{name:"wangergou",age:30,sex:"nan"}
	w1.work() //两种方法的写法均可
	w1.rest()

	w2 := &worker{name:"C++", age :16,sex:"male"}
	fmt.Printf("%T\n",w2)
	w2.work()
	w2.rest()
}

type worker struct { //当命名(常量,变量,类型,函数名,结构体等)以大写字母开头证明可以被外部包使用public 小写的话包内可见private
	name string
	age int
	sex string
}

//正常的函数声明 func getSum2(a, b int) int
func (w worker) work(){
	fmt.Println(w.name,"working......")
}

func (p *worker) rest(){
	fmt.Println(p.name,"resting......")
}

type Person struct {
	name string
	age int
}

type Student struct {
	Person //这里其实就是模拟了继承
	school string //如果匿名字段是结构体,那么这个结构体中的成员我们成为提升字段,可以不通过这个结构体直接访问
}