package main

import (
	"fmt"
	"time"
)

func main() {
	// sync.RWMutex{}
	// 当涉及到多go程时，c语言使用互斥量，上锁来保持资源同步
	// go语言也支持这种方式，但是go语言更好的解决方案是使用管道、通道
	// 使用通道不需要我们去加解锁
	// A 写 B 读
	// 读写的数量必须相同，不然会崩溃

	// 创建管道：创建一个装数字的管道
	// 不设定长度，则为无缓冲
	// make(chan int, 10) 可以设置缓冲区长度
	numChan := make(chan int) // 使用管道的使用一定要make，同map一样，否则是nil
	// strChan := make(chan string)

	// 创建两个go程，父亲写数据，儿子读数据
	go func() {
		for i := 0; i < 150; i++ {
			data := <-numChan
			fmt.Println("子go程1读取：", data)
		}
	}()

	go func() {
		for i := 100; i < 200; i++ {
			numChan <- i
			fmt.Println("++++子go程2写入：", i)
		}
	}()
	for i := 0; i < 50; i++ {
		numChan <- i
		fmt.Println("====这是主go程：写入数据：", i)
	}

	time.Sleep(5 * time.Second)
}
