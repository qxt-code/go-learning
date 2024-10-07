package src

import "fmt"

// go语言中通过首字母大小写控制public 和private
// 1. import ->如果包名不同，那么只有大写字母开头的才是public
// 2. 对于类里面的成员、方法，只有大写开头的才能在其他包中使用

type Human struct {
	Name   string
	Age    int
	Gender string
}

// 方法绑定
// 指针与非指针：指针可以修改调用的实例，不使用指针使用的是副本
// func (this *Person) Eat(){}
func (p Human) Eat() {
	fmt.Println(p.Name, "is eating")
}

// 嵌套
type Student1 struct {
	Hum    Human
	School string
	Score  float64
}

// 继承
// 虽然没有指定字段名字，但是会自动创建一个默认的同名字段
// 这是为了在子类中依然可以操作父类，因为子类、父类可能出现同名字段
type Teacher struct {
	Human   // 直接写类型，没有字段名
	Subject string
}
