package main

import "fmt"

func main() {
	/*
			位运算符中^不只是异或,也是按位取反

		位清空: &^
			a &^ b
			对于b上的每一个位
		如果为0 则取a对应位上的值
		如果为1 则取则取0
	*/
	a := 60
	b := 13
	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)

	c := a &^ b
	fmt.Printf("%d, %b\n", c, c)
}
