package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串是一些字节的集合
	// 可以理解为一个字节的序列 有对应的索引
	s1 := "Hello世界"
	s2 := `hello world`
	fmt.Println(s1, s2)

	fmt.Println(len(s1)) //中文字符占三个字节
	fmt.Println(len(s2))

	fmt.Println(s1[0]) // 显示72 表示ascll码值

	//字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)

	s4 := "abcdefg"
	slice2 := []byte(s4) //字符串可以和字节进行转换
	fmt.Println(slice2)
	fmt.Printf("%T\n", slice2) // []uint8

	//字符串是字节的集合 strings 包中都是字符串相关函数
	s11 := "helloworld"
	fmt.Println(strings.Contains(s11, "world"))
	fmt.Println(strings.Count(s11, "l")) //查找子串的数目

	s44 := "12,3,64,6,65,6,56,54564"
	ss2 := strings.Split(s44, ",")
	fmt.Println(ss2)

	s7 := "HELOlsdas ASDsdc as"
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))
	s8 := s7[0:5] // 没有substring类似的函数 直接使用切片即可
	fmt.Println(s8)

	// strconv包的使用
	ss1 := "true"
	bb1, err := strconv.ParseBool(ss1)
	if err != nil {
		fmt.Println(bb1)
		return
	}
	fmt.Printf("%T %t\n", bb1, bb1)
	bb2 := false
	d2 := strconv.FormatBool(bb2)
	fmt.Println(d2)

	d1 := "1000"
	i2, err := strconv.ParseInt(d1, 10,64)
	// 第三个参数
	/*
	The bitSize argument specifies the integer type that the result must fit into.
	Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64.
	If bitSize is below 0 or above 64, an error is returned
	 */
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(i2)

}
