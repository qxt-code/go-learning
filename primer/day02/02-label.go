package main

import "fmt"

func main() {
	// 标签 LABEL1
	// goto LABEL1 -->下次进入循环时，i不会保存之前的状态
	// break LABEL1  --> 直接跳出指定位置的循环，不会重新进入
	// continue LABEL1 -->跳到指定未知，但是会保存循环变量状态

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				// goto LABEL1
				// continue LABEL1
				break LABEL1
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	fmt.Println("over!")
}
