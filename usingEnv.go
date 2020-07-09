package main

import (
	"exercise/env"
	"fmt"
)

func main(){
	fmt.Println(tempenv.AbsoluteZeroC)
	c:= tempenv.FtoC(212.0)
	fmt.Println(float64(c))
}

