package main

import (
	"fmt"
	"os"
)

// 从命令行输入惨护士，在switch中进行处理

func main() {
	// C: argc, **argv
	// Go: os.Args ==>直接可以获取命令输入，是一个字符串切片
	cmds := os.Args

	// os.Args[0]程序名
	// os.Args[1]第一个参数
	for key, cmd := range cmds {
		fmt.Println("key : ", key, ", cmd :", cmd, ", cmds len: ", len(cmds))
	}

	if len(cmds) < 2 {
		fmt.Println("请正确输入参数！")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		// go的switch 默认加break
		// 如果想向下穿透的话，需要加上关键字：fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("dafault")
	}
}
