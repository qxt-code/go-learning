package main

import "fmt"

// 内存逃逸，栈上的变量返回后，会自动移动到堆上
func main() {
	p1 := testPtr1()
	fmt.Println("p1 :", *p1)
}

func testPtr1() *string {
	name := "Duke"
	p0 := &name
	fmt.Println("*p0:", *p0)

	city := "shenzhen"
	ptr := &city
	return ptr
}
