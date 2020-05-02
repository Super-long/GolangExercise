package main

import (
	"fmt"
	"math"
)

func main() {
	/*area1, err1 := cricleArea(-1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(area1)*/
	//----------------
	area2, err2 := rectArea(-5, 16)
	if err2 != nil {
		fmt.Println(err2) //当产生错误的时候通过断言去访问本来的错误类型,然后可以使用原类型的函数
		if err, ok := err2.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Println("长度小于0")
			}
			if err.widthNegative() {
				fmt.Println("宽度小于0")
			}
		}
		return
	}
	fmt.Println("面积为:", area2)
}

type areaError struct {
	msg    string
	radius float64
}

/*
type error interface { //错误的接口
	Error() string
}
*/
func (a *areaError) Error() string {
	return fmt.Sprintf("%f -> %s", a.radius, a.msg)
}

func cricleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"此半径非法", radius}
	} else {
		return 2 * math.Pi * radius, nil
	}
}

//----------------------------

type areaError2 struct {
	msg    string
	length float64
	width  float64
}

func (s *areaError2) Error() string {
	return s.msg
}

func (s *areaError2) lengthNegative() bool {
	return s.length < 0
}

func (s *areaError2) widthNegative() bool {
	return s.width < 0
}

func rectArea(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "长度小于0"
	}
	if width < 0 {
		if msg != "" {
			msg = "长度,宽度都小于0"
		} else {
			msg = "宽度小于0"
		}
	}
	if msg != "" {
		return 0, &areaError2{msg, length, width}
	} else {
		return length * width, nil
	}
}
