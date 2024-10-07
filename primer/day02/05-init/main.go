package main

import (
	_ "lesson/day02/05-init/sub"
	// . "lesson/12-import/sub" .代表使用函数时不需要加包名
)

func main() {
	// res := sub.Sub(20, 10)
	// fmt.Println(res)

	// 这个无法调用，是因为函数的首字母是小写的
	// 如果一个包里面的函数想对外提供访问权限，那么一定要首字母大写
	// 大写字母开头相当于：public
	// 小写字母开头相当于：private
	// add.add(10, 20)
}
