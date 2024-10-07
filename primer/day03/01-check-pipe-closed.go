package main

import "fmt"

func main() {
	numChan := make(chan int, 10)

	// 写
	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
			fmt.Println("写入：", i)
		}
		close(numChan)
	}()

	// 读
	// for v := range numChan {
	// 	fmt.Println("v: ", v)
	// }
	for {
		v, ok := <-numChan
		if !ok {
			fmt.Println("管道已经关闭")
			break
		}
		fmt.Println("v: ", v)
	}
	fmt.Println("OVER")
}
