package main

import "fmt"

func main() {
	// 其实也就是动态数组
	var slice1 []int // 不写长度就是切片
	fmt.Println(slice1)
	s3 := make([]int,3,8)
	fmt.Println(s3)
	fmt.Printf("%d %d\n", cap(s3), len(s3))

	s4 := make([]int,5,5)
	s4 = append(s4,5,6,4) //返回值如果不指定或者抛弃的话会报错
	s3 = append(s3, 1,21,5,6,5,4,5,6) //append还可以直接加一个slice
	s3 = append(s3, s4...)
	fmt.Println(s3)
	fmt.Printf("%d %d\n", cap(s3), len(s3)) //容量一次增加二倍

	a := []int{1,2,3,4,5,6,7,8,9,10}
	s1 := a[:5]
	s2 := a[3:8]
	s5 := a[5:]
	s6 := a[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",s1)

	s1 = append(s1, 1,1,1,1,1,1,1,1) //如果append以后小于容量修改原数组 否则的话重新建立一个数组
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2) //从slice衍生出来的其实都指向同一个底层的数组
	s22 := make([]int,0)
	for i :=1;i<5;i++{
		s22 = append(s22, i)
	}
	s32 := make([]int,8,8)
	//copy(s32, s22) //C++使用迭代器,这其实就是一个迭代器的封装,不过方便的多
	copy(s32[1:], s22[2:])
	fmt.Println("-----------------")
	fmt.Println(s22)
	fmt.Println(s32)


}
