package main

import "fmt"

type Student1 struct {
	Name  string
	Age   int
	Score int
}

// 方法定义 func (方法接收者) 方法名 (参数列表) (返回值列表){}

// 值对象
// 无需修改对象
// 引用类型  字符串  函数
// 调用时会按照其一个副本来执行调用
func (s Student1) SayHello1() {
	//fmt.Println(s.Name,": Hello")
	s.Name = "李四"
}

// 指针对象接收者
// 需要修改的数据
// 大对象地址操作
// 按照实际的值来调用执行
func (s *Student1) SayHello2() {
	//fmt.Println(s.Name,": Hello")
	s.Name = "李四"
}
func main() {
	stu := Student1{"张三", 18, 100}
	fmt.Println(stu)
	stu.SayHello1()
	fmt.Println(stu)
	stu.SayHello2()
	fmt.Println(stu)
}
