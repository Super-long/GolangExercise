package main

import "fmt"

// channel可以隐式转换为一个一个单工管道
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in{
		out <- x*x
	}
	close(out)
}

func printer(in <-chan int){
	for x := range in{
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

/*	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}*/

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
