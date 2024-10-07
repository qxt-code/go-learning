package main

import (
	"fmt"
	"time"
)

// 1. 当管道写满了，写阻塞
// 2. 当缓冲区读完了，读阻塞
// 3. 如果管道没有使用make分配空间，管道默认是nil
// 4. 从nil的管道读取、写入都会阻塞（不会崩溃）
// 5. 从一个已经close的管道读取数据时，会返回零值（不会崩溃）
// 6. 向一个已经close的管道写数据时，会崩溃
// 7. 关闭一个已经close的管道，程序会崩溃
// 8. 关闭管道的动作一定要放在写端，不应该放在读端
// 9. 读和写的次数一定要对等
// 		1. 在多个go程中：资源泄漏
//  	2. 在主go程中，程序崩溃（deadlock）

func main() {

	numsChan := make(chan int, 10)

	// 写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan <- i
			fmt.Println("写入：", i)
		}
		fmt.Println("数据写完毕，准备关闭管道")

		close(numsChan)
	}()

	// 遍历管道时，只返回一个值
	// for range不知道管道是否已经写完了，所以会一直在这里等待
	// ----解决方案：在写入端主动关闭，关闭管道时，for range会退出
	for v := range numsChan {
		fmt.Println("读取：", v)
	}

	for {
		fmt.Println("主go程正在死循环。。。。。。。。。。。。")
		time.Sleep(1 * time.Second)
	}
}
