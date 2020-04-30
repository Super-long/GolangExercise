package main

import "fmt"

func main() {
	/*
	基本数据类型
	bool
	数值
	字符串
	复合数据类型
	array,slice,map,function,pointer,struct,interface,channel(特有)
	 */
	fmt.Printf(`"asdsad" 不需要转义 类似于原生字符串\n`)

	var a int8
	a = 10

	var b int16
	b = int16(a)

	fmt.Println(b)

	f1 := 4.83
	var c int
	c = int(f1)
	fmt.Println(c)
}
