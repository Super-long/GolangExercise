package main

import "fmt"

/*
接口最大的作用就是解耦合

go接口最大的特点就是非侵入式的 我们不需要显式的实现接口,只需要正常的实现接口的方法即可,编译器会在使用的时候检测到
 */

func main() {
	p1 := Mouse{"nihao"}
	testInterface(p1)

	p2 := FlashDisk{"world"}
	testInterface(p2)

	var usb USB
	usb = p1
	usb.start()
	usb.end()
	//fmt.Println(usb.name)

}

// 定义接口
type USB interface {
	start() //开始工作
	end()   //结束工作
}

// 实现类
type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("鼠标开始工作啦")
}

func (m Mouse) end() {
	fmt.Println("鼠标结束工作啦")
}

type FlashDisk struct {
	name string
}

func (m FlashDisk) start() {
	fmt.Println("U盘开始工作啦")
}
//如果有一个位实现的话报如下信息
/*
 cannot use p2 (type FlashDisk) as type USB in argument to testInterface:
        FlashDisk does not implement USB (missing start method)

 */

func (m FlashDisk) end() {
	fmt.Println("U盘结束工作啦")
}

// 测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
