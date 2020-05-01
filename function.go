package main

import "fmt"

//num3 := 100 全局变量不能使用简短定义
var num3 int = 100

func main() {
	fmt.Println(num3)
	getAnySum()
	getAnySum(1, 2, 3)

	s1 := []int{1, 2, 3, 4, 5}
	getAnySum(s1...) //可变函数参数就相当于是一个变长数组
	//函数中涉及到一个值传递和引用传递
	/*
		A:值传递   值类型的数据默认都是值传递 处理下面三个以外都是
		B:引用传递 slice,map,channel
	*/
	fmt.Println(s1)
	fun2(s1)
	fmt.Println(s1)

	fmt.Println(getSum2(10, 15))
	fmt.Println(getSum3(10, 15))

	perimeter, area := rectangle(12.34, 163.56) //提到多返回值可以使用_丢弃一个返回值
	fmt.Println(perimeter, area)

	//-------------------

	defer fun1("hello")
	fmt.Println("123456")
	defer fun1("world") //引入一个外围函数的概念 就是当前所处函数 当外围函数结束以后执行的函数 类似于智能指针
	//其实就是在return之前把所有defer过的函数执行一遍
	fmt.Println("dowm defer")

	// defer的函数被加入一个栈中

	//-------------------
	a := 2
	fmt.Println(a)
	defer fun3(a) // 可以看到其实值还是不变的 证明其实值已经存进去了 类似于一个function
	a++
	fmt.Println("dowm defer")

	//-------------------
	//匿名函数 其实还是用作函数式编程
	/*
	https://www.bilibili.com/video/BV1Db411s7in?p=78
	1.作为另一个函数的参数
	2.作为另一个函数的返回值,可以形成闭包
	这里与C++不同的是闭包不涉及捕获,当有闭包结构的时候临时变量的值会自动延长
	 */
	fun4 := func() {
		fmt.Println("see hello!")
	}
	fun4()
}

func fun3(a int) {
	fmt.Println(a)
}

func fun1(s string) {
	fmt.Println(s)
}

func getSum(a int, b int) {

}
func getSum2(a, b int) int { //类型相同的时候这样也可 传入参数类型必须相同
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func getSum3(a, b int) (sum int) { //相当于先定义一个值为该类型默认值的变量 然后后面直接return就可以了
	for i := a; i <= b; i++ {
		sum += i
	}
	return
}

func getAnySum(arg ...int) { //可变参数类型
	sum := 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	fmt.Println(sum)
}

func fun2(s2 []int) {
	s2[0] = 100
}

func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}
