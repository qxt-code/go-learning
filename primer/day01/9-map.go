package main

import "fmt"

func main() {
	// 1.定义字典
	var idNames map[int]string // 定义一个map，此时不能直接赋值，需要分配空间

	// 2.分配空间，使用make
	idScore := make(map[int]float64)
	idNames = make(map[int]string, 10) //可以不指定大小，建议指定大小

	// 3. 定义时直接分配空间
	// idNames := make(map[int]string, 10)

	idNames[0] = "duke"
	idNames[1] = "lily"

	// 4.遍历
	for key, value := range idNames {
		fmt.Println(key, ", ", value)
	}

	// 5.如何确定一个key是否存在一个map中
	//在map中不存在越界问题，它认为所有的key也都是有效的，所以访问一个不存在的key不会崩溃，返回这个类型的零值
	// 零值： bool -> false， 数字-> 0, 字符串-> 空
	names9 := idNames[9]
	fmt.Println("names9:", names9)
	fmt.Println("idScore[100] : ", idScore[100])

	// 无法通过获取value判断一个key是否存在，因为存在的key的value也可能是零值
	value, ok := idNames[10] // 如果key=1存在，那么value既是key=1对应的值，ok返回true，反之返回零值，ok返回false
	if ok {
		fmt.Println("id=1存在,value为：", value)
	} else {
		fmt.Println("id=10不存在,value为：", value)
	}

	// 6.删除map中的元素
	// 使用自由函数delete来删除指定的key
	fmt.Println("idNames删除前：", idNames)
	delete(idNames, 1)
	fmt.Println("idNames删除后：", idNames)
	delete(idNames, 100)
	fmt.Println("idNames删除不存在元素：", idNames)

	// 并发时需要上锁，TODO
}
