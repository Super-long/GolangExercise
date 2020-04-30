package main

import "fmt"

func isHappy(n int) bool {
	mp := map[int]bool{}
	for ; n != 1 && !mp[n];n, mp[n] = step(n), true{}
	return n == 1
}

func step(n int) int{
	sum := 0
	for n > 0{
		sum += (n%10)*(n%10)
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
