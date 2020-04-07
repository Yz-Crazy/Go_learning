package main

import "fmt"

type People2 struct {
	Name string
	Age  int
}

// 建议使用指针对象接收者
func (p *People2) GetName() {
	fmt.Println("父类方法", p.Name)
}

type Student2 struct {
	People2 // 继承了People2（组合）
	//Name string
	Score int
}

// 同一个对象的方法名不能重名
func (s *Student2) GetName() {
	fmt.Println("子类方法", s.Name)
}

func main() {
	// 子类继承父类    子类可以使用父类结构体成员属性，可以使用父类方法
	// 子类和父类的方法名可以重名，叫做方法重写
	s := Student2{People2{"张三", 18}, 100}
	// 子类方法
	s.GetName()
	// 子类调用父类方法
	s.People2.GetName()

}
