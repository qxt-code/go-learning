package main

import (
	"fmt"
	"time"
)

// 单向通道：为了明确语义，一般用于函数参数

func main() {
	// 单向读通道：
	// var numChanReadOnly <- chan int
	// 单向写通道：
	// var numChanWriteOnly chan <- int

	// 生产者消费者模型
	// C:数组+锁， thread1 ： 写， thread2: 读
	// go ：goroutine + channel

	// 1. 在主函数中创建一个双向通道
	numChan := make(chan int, 5)
	// 2. 将numChan 传递给producer，负责生产
	go producer(numChan) // 双向通道可以赋值给同类型的单向通道， 单向不能转双向
	// 3. 将numChan 传递给consumer，负责消费
	go consumer(numChan)
	time.Sleep(1 * time.Second)
	fmt.Println("OVER")
}

func producer(out chan<- int) {
	for i := 0; i < 50; i++ {
		out <- i
		fmt.Println("写入：", i)
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("读取：", v)
	}
}
