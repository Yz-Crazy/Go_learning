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
}

func main() {
	testMap()
	testPointer()
}
