package main

import (
	"fmt"
	"time"
)

// go程：比线程轻量级的线程，一个go程只需要4-5KB内存
// 一个进程可以启动上千go程，而只能启动几十个线程
// 启动go程只需要在调用函数前加一个go关键字

func display(i int) {
	count := 1
	for {
		fmt.Println("====================>这是子go程:", i, "当前count:", count)
		count++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 自动子go程
	for i := 0; i < 3; i++ {
		go display(i)
	}

	// 主go程
	count := 1
	for {
		fmt.Println("这是主go程:", count)
		count++
		time.Sleep(2 * time.Second)
	}
}
