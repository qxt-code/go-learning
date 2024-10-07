package main

import (
	"fmt"
	"time"
)

// 当程序中有多个channel协同工作ch1、ch2，某一个时刻，ch1或ch2触发了，程序要做响应处理
// 使用select处理

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	// 启动一个go程，负责监听两个channel
	go func() {
		for {
			fmt.Println("监听中")
			select {
			case data1 := <-chan1:
				fmt.Println("从chan1读取：", data1)
			case data2 := <-chan2:
				fmt.Println("从chan2读取：", data2)
			default:
				fmt.Println("select default")
				time.Sleep(time.Second)
			}
		}
	}()
	// 启动go1 写chan1
	go func() {
		for i := 0; i < 50; i++ {
			chan1 <- i
			time.Sleep(1 * time.Second / 2)
		}
	}()
	// 启动go2 写chan2
	go func() {
		for i := 0; i < 50; i++ {
			chan2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		fmt.Println("OVER")
		time.Sleep(5 * time.Second)
	}

}
