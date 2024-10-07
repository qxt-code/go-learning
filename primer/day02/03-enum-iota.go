package main

import "fmt"

// 在go语言中没有枚举类型，但是我们可以用const + iota(常量累加器)来进行模拟
// 模拟一个一周的枚举
const (
	MONDAY    = iota
	TUESDAY   = iota
	WEDNESDAY = iota
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
	M, N = iota, iota
)

const (
	JANU = iota + 1
	FER
)

/*
1. iota是常量组计数器
2. iota从0开始，每换行递增1
3. 常量组如果不赋值，默认与上一行表达式相同
4. 如果同一行出现两个iota，那么两个iota的值相同
5. 每个常量组的iota是独立的
*/

func main() {

	fmt.Println("打印周：")
	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(M, N)

	fmt.Println("打印月：")
	fmt.Println(JANU)
	fmt.Println(FER)

	// 可以使用变量组统一定义变量
	// var {
	// 	number int
	// 	name string
	// 	flag bool
	// }
}
