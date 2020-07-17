package main

import (
	"fmt"
	"reflect"
)

type replace int64

func main() {
	/*	一个获取可寻址value的通用方法
		x := 2
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(2)
		c := reflect.ValueOf(&x)
		d := c.Elem()
		// false false false true
		fmt.Println(a.CanAddr(), b.CanAddr(), c.CanAddr(), d.CanAddr())*/

	/*	获取反射值原指针的一个通用方法,不过需要知道变量类型
		x := 2
		d := reflect.ValueOf(&x).Elem()		// 这是一种通用方法,可以返回一个可寻址的reflect.value
		fmt.Println(d.Type().Name())
		px := d.Addr().Interface().(*int)	// 这样可以修改x的值
		*px = 3
		fmt.Println(x)*/

	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)
	fmt.Println(x)
	rx.Set(reflect.ValueOf(3))
	fmt.Println(x)

	// 就算是底层类型为带符号整数类型的命名类型也可以转换
	var v replace
	v = 5
	rv := reflect.ValueOf(&v).Elem()
	rv.SetInt(6)
	fmt.Println(v)

	// rx.SetString("hello") // panic

	// 接口的话只能使用set
	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	ry.Set(reflect.ValueOf(3))
	fmt.Println(y.(int))

	ry.Set(reflect.ValueOf("hello"))
	fmt.Println(y.(string))

	//ry.SetInt(5) // panic
}
