package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Intn(45)
	fmt.Println(num1)
	rand.Seed(time.Now().UnixNano()) //随机种子需要一个64位整数
	for i := 0; i < 10; i++ {
		num1 := rand.Intn(10)
		fmt.Println(num1)
	}
	var arr1 [10]int
	arr1[0] = 1
	arr1[2] = 2
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1)) //数组两个是相同的

	var b = [5]int{1, 2, 3, 4}
	fmt.Println(b)
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	for index,value := range arr1{ //不需要下划线的话可以舍去
		fmt.Printf("index is %d, value is %d\n",index,value)
	}
	for _,value := range arr1{ //
		fmt.Printf("value is %d\n",value)
	}

	var arr3 [5][5]int
	arr3[0][0] = 2
	arr4:=[3][4] int{{1,2,3},{2,3,4},{4,5,6}}
	for _,value := range arr4{
		for _,vv := range value {
			fmt.Print(vv)
		}
		fmt.Println()
	}
	fmt.Println(arr3[0][0])
}
