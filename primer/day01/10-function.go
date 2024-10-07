package main

import "fmt"

// 1. 函数返回值在参数列表之后
// 2. 如果有多个返回值，需要使用圆括号包裹，多个参数之间使用逗号分割
func testfunc1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

// 连续的同类型函数可以只写一个类型名
func testfunc2(a, b int, c string) (res int, str string, bl bool) {
	var i, j, k int
	i = 1
	j = 2
	k = 3
	fmt.Println(i, j, k)

	// 直接使用返回值的变量名参与运算
	res = a + b
	str = c
	bl = true

	// 当返回值有名字时，可以直接return
	return
	// return a + b, c, true
}

// 如果返回值只有一个参数，并且没有名字，那么不需要加圆括号
func testfunc3() int {
	return 1
}

func main() {
	v1, s1, _ := testfunc1(10, 20, "hello")
	fmt.Println("v1:", v1, "s1:", s1)

	v3, s3, _ := testfunc2(20, 30, "hello")
	fmt.Println("v3: ", v3, " s3: ", s3)
}
