package main

import "fmt"

func main() {
	var num int
	num = 30;
	fmt.Printf("num的值是 : %d\n", num)
	var num2 int = 15;
	fmt.Printf("num2的只是%d\n", num2)
	var name = "nihaoya"
	fmt.Printf("类型是 : %T ;值是 %s\n ", name, name);
	num3 := 30 //:=代表定义变量 左边至少有一个新的 且全局变量不支持简短定义
	fmt.Printf("num3的只是%d\n", num3)

	var a,b,c int
	a = 1
	b = 1
	c = 1
	fmt.Println(a,b,c)
	var (
		studentName = "lixiaohau"
		age = 15
		sex = "nv"
	)
	fmt.Printf("%s %d %s\n",studentName, age ,sex)

	var m int //未赋值默认为零
	var s string //一个空的字符串
	fmt.Println(m, s)
	var s2 []int //切片[]
	fmt.Println(s2)
	fmt.Println(s2 == nil)
}