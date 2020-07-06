package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime1 := time.Now().UnixNano()
	s, sep := "", ""
	for index, args := range os.Args[1:] {
		s += sep + args
		sep = " "
		fmt.Println(index)
	}
	fmt.Println(s)
	endTime1 := time.Now().UnixNano()

	startTime2 := time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[0:], " "))
	endTime2 := time.Now().UnixNano()

	fmt.Println(endTime1 - startTime1)
	fmt.Println(endTime2 - startTime2)
}
