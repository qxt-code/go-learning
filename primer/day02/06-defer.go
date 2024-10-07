package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. 延迟，修饰语句、函数，在栈退出时执行
	// 2. 一般用于做资源清理工作
	// 3. 解锁、关闭文件
	// 4. 在同一个函数中多调用defer，执行时类似于栈的机制，后入先出
	readFile("data.txt")

}

func readFile(filename string) {
	// go语言一般将错误码作为最后一个参数
	// err一般Nil代表没有错误
	f1, err := os.Open(filename)
	// defer f1.Close()

	// 创建匿名函数并调用
	defer func() {
		fmt.Println("准备关闭文件")
		f1.Close()
	}()

	if err != nil {
		fmt.Println("open error")
		return
	}
	buf := make([]byte, 1024)
	n, err := f1.Read(buf)
	fmt.Println("读取长度: ", n)
	fmt.Println("读取的内容：", string(buf))

}
