package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		var x int
		var y float32
		var z string
		fmt.Scan(&x, &y, &z)
		fmt.Println(x,y,z)
	*/
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n') //以'\n'结尾

	var s2 int
	fmt.Scan(&s2)
	if s2 > 16 {
		fmt.Println(s2)
	}
	fmt.Println(s1)

	if num := 4; num > 0 {
		fmt.Println(num)
	}

	switch s2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("four")
	}
	//score:=78 //只输出一个值 代表在匹配到第一个的时候直接退出
	switch score := 78; {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60 && score < 80:
		fmt.Println("优秀")
	case score >= 60 && score < 90:
		fmt.Println("优秀啊")
	}

	switch temp := 15; temp {
	case 15:
		fmt.Println("one")
		fallthrough //使得下一个case不需要判断即可直接执行
	case 25:
		fmt.Println("two")

	}

	for i := 0; i < 5; i++ {
		fmt.Println("hello world")
	}
	i := 1
	for i <= 5 { //省略第一个和第三个表达式
		fmt.Println("nihao")
		i++
	}
	/*for{ // 如果省略表达式2 相当于while true
		fmt.Println("hhhhh")
	}*/

	Temp := 10
LOOP:
	for Temp < 20 {
		if Temp == 15 {
			Temp++
			goto LOOP
		}
		fmt.Println(Temp)
		Temp++
	}
}
