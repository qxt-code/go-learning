package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	gender string // 注意，小写
}

func main() {
	// 结构体===>字符串===>编码
	// ====>字符串===>结构体====

	lily := Student{
		Id:     1,
		Name:   "Lily",
		Age:    20,
		gender: "female",
	}
	// 编码（序列化）， 结构体===>字符串
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))

	// 对端接收到数据
	// 反序列化（解码）： 字符串====>结构体
	var lily_copy Student
	if err := json.Unmarshal([]byte(encodeInfo), &lily_copy); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	fmt.Println("name:", lily_copy.Name)
	fmt.Println("gender:", lily_copy.gender)
	fmt.Println("Age:", lily_copy.Age)
	fmt.Println("Id:", lily_copy.Id)
}
