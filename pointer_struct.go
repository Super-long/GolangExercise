package main

import "fmt"

type studnt struct {
	name string
	age int
	sex string
	address string
}

func main() {
	var p1 *int
	fmt.Println(p1)
	a := 100
	p1 = &a
	fmt.Println(*p1)

	//------------------- 数组指针

	arr1 := [4]int{1,2,3,4}
	fmt.Println(arr1)

	var p2 *[4]int // 指针数组 var p2 [4]*int
	p2 = &arr1
	fmt.Println(p2)
	fmt.Println(&arr1)
	fmt.Println(&p2)

	(*p2)[0] = 100
	fmt.Println(arr1)

	//--------------结构体 四种初始化方式 其实还可以用new
	var one studnt
	fmt.Println(one) //{ 0  } 全部都是默认值
	one.address = "shanxi"
	one.age = 18
	one.name = "lizahol"
	one.sex = "male"
	fmt.Println(one)

	two := studnt{}
	two.address = "shanxi"
	two.age = 18
	two.name = "zhangsiyao"
	two.sex = "female"
	fmt.Println(two)

	three := studnt{name:"zahngyao", age:20, sex:"female", address:"shanxi"}
	fmt.Println(three)

	four := studnt{"lizhaol",48,"male","feizhou"}
	fmt.Println(four)

	pp1 := new (studnt)
	fmt.Printf("%T\n", pp1)
	fmt.Println(pp1)
	pp1.address = "shanxi"
	pp1.age = 18
	pp1.name = "lizahol"
	pp1.sex = "male"
	fmt.Printf("%T\n", pp1)
	fmt.Println(pp1)

	//匿名字段
	s2 := struct {
		name string
		age int
	}{
		name :"zhangsiyao",
		age : 18,
	}
	fmt.Println(s2)

	w2 := Worker{"lixiaohua", 32}
	fmt.Println(w2)
	fmt.Println(w2.string) //说实话意义不大
}

type Worker struct {
	string
	int //匿名字段 默认使用数据类型作为名字 显然局限性较大
}