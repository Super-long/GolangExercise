package main

import "fmt"

func main() {
	const b string = "abc"
	const a = "abdc"
	fmt.Println(b, a) // 变量声明以后不使用会报错,但是常量不会
	const (
		aa int = 100
		ba     //如果没有赋值的话默认和上一个是类似的
	)
	fmt.Printf("%T, %d\n", aa, aa)
	fmt.Printf("%T, %d\n", ba, ba)

	// 可以使用常量组当做一个枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

	/*
	iota: 每当定义一个const iota初始值为0
	 */
	const (
		ac = iota
		bc = iota
		cc = iota
	)
	fmt.Println(ac)
	fmt.Println(bc)
	fmt.Println(cc)

	const (
		d = iota
		e
	)
	fmt.Println(d)
	fmt.Println(e)
}
