package main

import (
	"fmt"
	"math"
)

func main() {
	var ti triangle = triangle{3, 4, 5}
	fmt.Println(ti.perimeter())
	fmt.Println(ti.area())

	var ci cricle = cricle{4}
	fmt.Println(ci.perimeter())
	fmt.Println(ci.area())

	var s1 shape
	s1 = ti
	fmt.Println(s1.perimeter())
	fmt.Println(s1.area())

	var s2 shape
	s2 = ci
	fmt.Println(s2.perimeter())
	fmt.Println(s2.area())

	getType(ti)
	getType(ci)

	var s3 *triangle = &triangle{3,4,5}
	getType(s3)

	getType2(ti)
	getType2(ci)
	getType2(s3)

}

func getType(s shape) { //这里其实发生了两次拷贝　一次是这一行　还有if后的ins
	// predicate
	if ins, ok := s.(triangle); ok { //如果是的话 返回值赋给ins的实例 话说这个语句挺奇怪的 ok这个变量是什么情况
		fmt.Printf("是三角形 %.2f %.2f %.2f\n", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(cricle); ok {
		fmt.Printf("圆形　%.2f\n", ins.radius)
	} else if ins,ok := s.(*triangle);ok {
		fmt.Printf("%T %p\n",ins,&ins)
	} else {
		fmt.Printf("我也不知道了\n")
	}
}

func getType2(s shape){
	switch s.(type) {
	case triangle:
		fmt.Println("这是三角形\n")
	case cricle:
		fmt.Println("这是圆形\n")
	default:
		fmt.Println("什么也不是")

	}
}

func testShape(s shape) {
	fmt.Println(s.area(), s.perimeter())
}

type shape interface {
	perimeter() float64 // 周长
	area() float64      // 面积
}

type triangle struct {
	a, b, c float64
}

func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t triangle) area() float64 {
	p := t.perimeter() / 2
	s := math.Sqrt(p * (p - t.b) * (p - t.b) * (p - t.c))
	return s
}

type cricle struct {
	radius float64
}

func (c cricle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (c cricle) area() float64 {
	return math.Pi * c.radius * c.radius
}
