package main

import "fmt"

// C++使用纯虚函数代替接口
// go有专门的关键字interface
// interface不仅是实现接口，还可以接受任意类型

func main() {
	// func Println(a ..interface{}) (n int, err error)

	// 定义三个接口类型
	var i, j, k interface{}
	names := []string{"Duke", "Lily"}
	i = names
	fmt.Println("i代表切片数组：", i)

	age := 20
	j = age
	fmt.Println("j代表数字：", j)

	str := "hello"
	k = str
	fmt.Println("k代表字符串：", k)

	// 我们现在只知道k是interface，但是不能够明确指导它代表的类型
	kvalue, ok := k.(int)
	if !ok {
		fmt.Println("k不是Int")
	} else {
		fmt.Println("k是int ,值为：", kvalue)
	}

	//最常用的场景：把interface当成一个函数的参数，（类似于print），使用switch
	//根据不同类型，做相应的逻辑

	//创建一个具有三个接口类型的切片
	array := make([]interface{}, 4)
	array[0] = 1
	array[1] = "Hello world"
	array[2] = true
	array[3] = 3.14
	for _, value := range array {
		switch v := value.(type) {
		case int:
			fmt.Printf("当前类型为int，内容为:%d\n", v)
		case string:
			fmt.Printf("当前类型我string,内容为: %s\n", v)
		case bool:
			fmt.Printf("当前类型为bool,内容为：%v\n", v) // %v可以自动推导
		default:
			fmt.Println("其他类型")
		}
	}
}
