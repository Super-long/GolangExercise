package main

import (
	"exercise/env"
	"flag"
	"fmt"
)
// 使用falg.Value解析命令行信息

type celsiusFlag struct {
	tempenv.Celsius
}

func (f *celsiusFlag) Set(s string) error{
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C" :
		f.Celsius = tempenv.Celsius(value)
		return nil
	case "F":
		f.Celsius = tempenv.FtoC(tempenv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invaild temperature %q", s)
}

func exchange(name string, value tempenv.Celsius, usage string) *tempenv.Celsius{
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = exchange("temp", 20.0, "the temperature")

func main(){
	flag.Parse()
	fmt.Println(*temp)
}

