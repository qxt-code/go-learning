package main

import "fmt"

func main() {
	names := []string{"beijing", "shanghai", "guagnzhou", "shenzhen"}
	for i, v := range names {
		fmt.Println("i: ", i, " v: ", v)
	}

	//1.追加数据
	fmt.Println("追加前长度： ", len(names), " ，容量： ", cap(names))
	names = append(names, "hainan")
	fmt.Println("names: ", names)
	fmt.Println("追加后长度： ", len(names), " ，容量： ", cap(names))

	//2.切片有长度len()和容量cap()
	nums := []int{}
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len: ", len(nums), ", cap: ", cap(nums))
	}

}
