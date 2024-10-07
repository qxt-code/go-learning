package main

import (
	"fmt"
	"time"
)

func main() {
	// numChan := make(chan int, 10)
	// 1. 当缓冲写满时，写阻塞
	// 2. 当缓冲区读取完毕，读阻塞
	// 3. 如果管道没有使用make分配空间，那么管道默认是nil的，读取、写入都会阻塞
	// 4. 对于一个管道，读与写的次数必须对等，

	var names chan string
	names = make(chan string, 10)

	go func() {
		fmt.Println("names: ", <-names)
	}()

	names <- "hello" // names是nil的，写操作会阻塞在这里
	time.Sleep(1 * time.Second)

	numsChan := make(chan int, 10)

	// 写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan <- i
			fmt.Println("写入：", i)
		}
	}()
	// 读，当主程序被管道阻塞时，那么程序将锁死崩溃
	// 要求我们一定要读写一致
	// 1. 当管道读写次数不一致时
	// 		1. 如果阻塞在主go程，那么程序会崩溃
	//  	2. 如果阻塞在子go程，那么内存会泄漏
	go func() {
		for i := 0; i < 60; i++ {
			data := <-numsChan
			fmt.Println("读取：", data)
		}
	}()

	for {
		fmt.Println("主go程正在死循环。。。。。。。。。。。。")
		time.Sleep(1 * time.Second)
	}
}
