package main

import "fmt"

func main() {
	names := [7]string{"beijing", "shanghai", "guagnzhou", "shenzhen", "luoyang", "nanjing", "qinhuangdao"}

	// 基于names创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]

	// 切片可以基于一个数组灵活地创建新的数组
	names2 := names[0:3]
	fmt.Println(names2)

	// 切片是对原数组的引用
	fmt.Println(&names[0], &names2[0])

	// 1.从0号元素开始截取， 那么冒号左边的数字可以省略
	names3 := names[:5]
	fmt.Println(names3)

	// 2.截取到最后一个元素，那么冒号后边的数字可以省略
	names4 := names[5:]
	fmt.Println(names4)

	// 3.如果想全部使用，两边都可以省略
	names5 := names[:]
	fmt.Println(names5)

	// 4.切片也可以用于获取子串
	sub1 := "helloworld"[5:7]
	fmt.Println(sub1)

	// 5.可以在创建空切片的时候，明确指定容量,第三个参数非必须，默认与长度相同
	// 长度范围内的空间才能用索引直接访问，长度范围外需要先append
	str2 := make([]string, 10, 20)
	fmt.Println("len: ", len(str2), ", cap : ", cap(str2))
	for i := 0; i < 10; i++ {
		str2[i] = "world"
	}
	fmt.Println(str2)

	// 6.copy 切片深拷贝
	namesCopy := make([]string, len(names))
	// copy接收的参数是切片
	copy(namesCopy, names[:])
	fmt.Println(namesCopy)
	namesCopy[0] = "xianggang"
	fmt.Println(names)
	fmt.Println(namesCopy)

}
