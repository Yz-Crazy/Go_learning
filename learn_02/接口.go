package main

import "fmt"

// type 接口变量 interface{方法声明}

type Animal interface {
	GetName() // 功能
}

// 方法
type People1 struct {
	Name  string
	Age   int
	Score int
}

// 接口功能的实现（大致意思是用一个方法实现这个接口的功能）
func (p *People1) GetName() {
	fmt.Printf("大家好，我叫%s,今年%d，考试成绩%d\n", p.Name, p.Age, p.Score)
}

type People3 struct {
	Id string
	Lv int
}

func (p *People3) GetName() {
	fmt.Printf("我叫%s，我的游戏等级是%d\n", p.Id, p.Lv)
}

// 函数  多态  将接口作为函数参数
func GetName(a Animal) {
	a.GetName()
}
func main() {
	// 定义接口类型
	var a Animal
	s := People1{"张三", 18, 100}
	p := People3{"幽灵", 120}
	// 必须将对象的地址赋值给接口类型变量
	a = &s
	a.GetName()
	a = &p
	a.GetName()

	GetName(&s)
	GetName(&p)
}
