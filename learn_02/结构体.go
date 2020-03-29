package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Student struct {
	//
	name string // 16    字节数
	sex  byte   // 1    7
	age  int    // 8
	addr string // 16
	aa   int32
	bb   int32
}

// 空结构体
type Null struct{}

// 结构体标签
type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 匿名字段
type people struct {
	Name string
	Age  int
}

// 结构体嵌套结构体（结构体组合）
type student struct {
	people
	Name  string
	Score int
}

func main() {
	var s Student
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s.name)
	fmt.Printf("%p\n", &s.sex)
	fmt.Printf("%p\n", &s.age)
	fmt.Printf("%p\n", &s.addr)
	fmt.Printf("%p\n", &s.aa)
	fmt.Printf("%p\n", &s.bb)
	// 计算结构体在内存中占用的字节大小
	fmt.Println(unsafe.Sizeof(s))

	// 空结构体
	// 只要是空结构体，他们都指向同一块内存空间
	// 在特殊的channel中使用，不可以写入数据，只有close关闭才能进行输出操作
	// 使用空 struct 是对内存更友好的打开方式，在 Go 源代码中针对 空struct 类数据内存申请部分，返回地址都是一个固定的地址，那么就避免了可能的内存滥用
	a := struct{}{}
	b := Null{}
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

	// 通过反射获取tag的设置
	to := reflect.TypeOf(People{})
	fmt.Printf("%T\n", to) // 查看类型
	for i := 0; i < to.NumField(); i++ {
		field := to.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}
	// 结构体嵌套结构体， 匿名字段的使用

	p := student{people{"张三", 18}, "张飞", 100}
	fmt.Println(p)
	fmt.Println(p.Name) // 先找自己结构体下的字段，如果没有，才去子结构体下去找
	fmt.Println(p.people.Name)

}
