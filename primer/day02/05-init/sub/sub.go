package sub

import "fmt"

// 1. init函数没有参数，没有返回值，原型固定如下
// 2. 一个包中包含多个init时，调用顺序是不确定的(同一个包中的多个文件中的init都会调用)
// 3. init函数不允许用户显式调用
// 4. 有的时候引用一个包，可能只想用里面的init函数（mysql的init堆驱动进行初始化）
// 但是不想使用里面的其他函数，为了防止编译器报错，可以使用_形式来处理
// 应用场景：configManager
func init() {
	fmt.Println("this is first init() in package sub")
}

func init() {
	fmt.Println("this is second init() in package sub")
}

// go语言中，同一层级目录下，不允许出现多个包名

func Sub(a, b int) int {
	test4() // 由于在同一个包下面，所以可以使用，并且不需要加包的路径
	return a - b
}
