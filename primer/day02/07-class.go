package main

import "fmt"

// go用结构体模拟类

type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

// 方法绑定
// 指针与非指针：指针可以修改调用的实例，不使用指针使用的是副本
// func (this *Person) Eat(){}
func (p Person) Eat() {
	fmt.Println(p.name, "Person is eating")
}

func main() {
	lily := Person{
		name: "Lily",
	}
	lily.Eat()
}
