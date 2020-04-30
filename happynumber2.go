package main

import (
	"fmt"
)

func isHappy1(n int) bool {
	slow := n
	quick := step(n)
	for slow != quick && quick != 1 {
		slow = step(slow)
		quick = step(step(quick))
	}
	return quick == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy1(19))
}
