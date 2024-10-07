package main

import (
	"09-priveledge/09-priveledge/src"
	"fmt"
)

func main() {
	s1 := src.Student1{
		Hum: src.Human{
			Name:   "Lily",
			Age:    18,
			Gender: "female",
		},
		School: "yizhong",
		Score:  80,
	}
	fmt.Println(s1.Hum.Name, s1.School)

	t1 := src.Teacher{}
	t1.Name = "Bob"
	t1.Age = 30
	t1.Gender = "male"
	t1.Subject = "gaoshu"
	fmt.Println(t1)
	t1.Eat()

	fmt.Println(t1.Name, t1.Subject)
}
