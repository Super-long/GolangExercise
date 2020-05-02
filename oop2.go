package main

import "fmt"

func main() {
	p1 := person{name : "lisa",age:16}
	fmt.Println(p1)
	p1.eat()

	p2 := student{p1, "youdaindaxue"}
	fmt.Println(p2.name)
	fmt.Println(p2.age) //子类对象可以访问父类的对象
	p2.eat() // 子类对象可以访问父类的方法
}

type person struct {
	name string
	age int
}

type student struct {
	person
	School string
}

func (p person) eat(){
	fmt.Println("吃窝窝头")
}

func (s student) study(){
	fmt.Println("新增的方法 学习啦")
}

func (s student) eat(){ //这其实就是多态
	fmt.Println("子类的方法 吃炸鸡")
}