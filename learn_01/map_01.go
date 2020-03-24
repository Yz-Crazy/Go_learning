package main

import (
	"fmt"
	"unsafe"
)

// map
func testMap() {
	// map 的key的类型 必须支持 == != 运算符的操作，不可以是函数 map 切片
	// map是无序的   在做查找时，速度很快
	// var maps map[int]string
	m := make(map[int]string)
	m[1001] = "法师"
	m[1002] = "战士"
	fmt.Println(m)
	// 删除元素
	//delete(m,1001)
	//fmt.Println(m)

	// 计算数据类型在内存中占的字节大小
	fmt.Println(unsafe.Sizeof(m))
	// map 是维护了一个hash表

}

// 指针
func testPointer() {
	// 所有指针类型在32位操作系统下占4个字节
	// 所有指针类型在64位操作系统下占8个字节
	// 内存地址是一个无符号十六进制整型数据

	// 通过指针可以修改指向变量地址的值
	// & 取地址运算符
	// * 取值运算符

	var a int = 10
	fmt.Println(&a)
	var p *int = &a
	fmt.Println(*p)
	*p = 123
	fmt.Println(*p)
	// 打印地址
	fmt.Println(p)
	// 内存地址编号 0 到 255 不允许用户读写
	//var c *int = nil
	//*c = 100
	//fmt.Println(c)
	//

	// 指针变量
	// var 指针变量 *数据类型=&变量
	// 指针变量除了有正确指向，还可以通过new()函数来指向
	// new(数据类型)
	// new() 函数的作用就是动态分配空间，不需要关心该空间的释放，Go 语言会自动释放
	var s *int
	// 创建一个int大小的内存空间，返回值为 *int
	s = new(int)
	*s = 123
	// 打印值
	fmt.Println(*s)
	// 打印地址
	fmt.Println(s)

}

func main() {
	testMap()
	testPointer()
}
