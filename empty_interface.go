package main

import "fmt"

// 空接口
/*
fmt包下很多都是这样
 Println(a ...interface{}) (n int, err error)
 */
func main() {
	var a1 A = Cat{"花猫"}
	var a2 A = person{"erghou",30}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test1(a1)
	test1(a2)
	test1(a3)
	test1(a4)

	map1 := make(map[string]interface{})
	map1["siyao"] = "xiaohua"
	map1["sadasd"] = "30"
	map1["zhalong"] = person{"jerry",30}
	fmt.Println(map1)

	slice1 := make([]interface{},0,10)
	slice1 = append(slice1,a1,a2,a3,300,"nihaoya") //虽然可以这样写 但是意义不大
	fmt.Println(slice1)
}

// 使用空接口可以接收所有的类型
func test1(a A)  {
	fmt.Println(a)
}

func test2(a interface{}){ //和上面的写法一样 都是实现空接口
	fmt.Println(a)
}

type A interface {

}

type Cat struct {
	color string
}

type person struct {
	name string
	age int
}