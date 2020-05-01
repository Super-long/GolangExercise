package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 map[int]string //只有声明没有初始化 是一个nil 没办法赋值
	var map2 = make(map[int]string)
	var map3 = map[string]int{"G3": 80, "Python": 87, "Html": 93}
	fmt.Println(map1, map2, map3)
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	map1 = make(map[int]string)
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	map1[1] = "lihzaolong"
	map1[2] = "zhangsiyao"
	map1[3] = "wangergou"
	fmt.Println(map1)
	fmt.Println(map1[3])
	fmt.Println(map1[100]) //这里是不存在的 但是仍会返回值 那么我们如何判断返回的值是空还是本身存的就是空呢
	value, ok := map1[100]
	if ok{
		fmt.Printf("存在 %s\n",value)
	}else {
		fmt.Println("不存在")
	}

	// 删除某个元素
	delete(map1,3)
	fmt.Println(map1)
	delete(map1,30)
	fmt.Println(map1)

	// 遍历
	/*
	使用 for range
	数组,slice :index value
	map: key value
	 */
	map4 := make(map[int]string)
	map4[1] = "one"
	map4[8] = "two"
	map4[3] = "three"
	map4[6] = "four"
	map4[5] = "five" //不是有序的
	for k,v := range map4{
		fmt.Println(k,v)
	}
	keys := make([]int,0,len(map4))
	for k,_ := range map4{
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys) // sort包中的方法
	fmt.Println(keys)

	s1 := []string{"Apple","Windows","Drange","abc","acd","wangergou"}
	sort.Strings(s1)
	fmt.Println(s1) //排序是大小写敏感的

	map5 := make(map[string]string)
	map5["name"] = "王二狗"
	map5["age"] = "38"
	map5["sex"] = "male"
	fmt.Println(map5)

	slice1 := make([]map[string]string,0,3) //将map放到slice里面
	slice1 = append(slice1, map5)

	for i,val := range slice1{
		fmt.Printf("第%d是 :",i)
		fmt.Println(val)
	}
}
