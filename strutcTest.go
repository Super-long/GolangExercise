package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type P struct {
	X, Y int
}

type Wheel struct{
	//P
	Circle
	Spokes int
}

func main(){
	var w Wheel
	w.X = 8
	w.Y = 8
	//w.P.Y = 8
	w.Spokes = 20
	w.Radius = 5
	fmt.Printf("%#v\n", w)
}