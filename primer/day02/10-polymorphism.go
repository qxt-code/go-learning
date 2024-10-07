package main

import "fmt"

// go语言的多态不需要继承，只需要实现相同的接口
// 人类的武器发起攻击，不同等级的子弹效果不同
// 每个接口都必须实现
// 定义接口
type IAttack interface {
	// 接口的函数可以有多个，但是只能有原型，不能有实现
	Attack()
}

// 低等级
type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是: ", a.name, ",等级为: ", a.level, "造成1点伤害")
}

// 高等级
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是: ", a.name, ",等级为: ", a.level, "造成999伤害")
}

// 定义一个多态的通用接口，传入不同的对象，调用相同的方法，实现不同的效果
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	var player IAttack

	lowLevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}

	highLevel := HumanHighLevel{
		name:  "Dave",
		level: 999,
	}
	lowLevel.Attack()

	// 对player赋值为lowLevel
	player = &lowLevel
	player.Attack()

	player = &highLevel
	player.Attack()

	fmt.Println("多态。。。。。。。。。。。。。")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
