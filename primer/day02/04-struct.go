package main

import "fmt"

// 类似typedef
type MyInt int

//go语言使用tyupe + struct定义结构体

type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt

	i, j = 10, 20
	fmt.Println(i, j)

	lily := Student{
		name:   "Lily",
		age:    20,
		gender: "female",
		score:  80, // 最后一个元素必须加逗号
	}
	fmt.Println(lily.name, lily.age, lily.gender, lily.score)

	s1 := &lily
	fmt.Println(s1.name, s1.age, s1.gender, s1.score)

	// 部分赋值必须指定key
	// 全部赋值可以省略字段名
	// duke := Student{
	// 	name: "Duke",
	// 	age:  28,
	// }
	duke := Student{
		"Duke",
		28,
		"male",
		99,
	}
	fmt.Println(duke)
}
