// package exercise
package main

import "fmt"

func main() {
	fmt.Println("HelloWorld")
	age := 30
	if age >= 30 {
		fmt.Println("你已经成年了")
	}

	list := []string {};
	list =  append(list, "nihao", "world")
	fmt.Println(fmt.Sprintf("%q", list))
}
