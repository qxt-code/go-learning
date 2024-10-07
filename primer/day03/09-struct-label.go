package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`                 // ===>使用Json编码时，不参与编码
	Subject string `json:"Subject_name"`      // 编码时，字段名会编码为Subject_name
	Age     int    `json:"age,string"`        // 改名+改类型, age与string中间不能有空格
	Address string `json:"address,omitempty"` // 如果是空值，编码时忽略
	gender  string // 注意，小写
}

func main() {
	t1 := Teacher{
		Name:    "Duke",
		Subject: "Golang",
		Age:     18,
		Address: "beijing",
		gender:  "male",
	}

	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal((&t1))

	fmt.Println("encodeInof: ", string(encodeInfo))
}
