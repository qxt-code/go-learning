package main

import (
	"fmt"
	"runtime"
	"time"
)

// GOEXIT 提前退出当前go程
// return 返回当前函数
// exit 退出当前进程
func main() {
	go func() {
		func() {
			fmt.Println("这是子go程内部的函数")
			// return
			// os.Exit(-1)
			runtime.Goexit()
		}()
		fmt.Println("子go程结束")
	}()

	fmt.Println("这是主go 程")
	time.Sleep(5 * time.Second)
	fmt.Println("OVER")
}
