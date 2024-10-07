package sub

// go语言中，同一层级目录下，不允许出现多个包名

func Sub(a, b int) int {
	test4() // 由于在同一个包下面，所以可以使用，并且不需要加包的路径
	return a - b
}
