package main

import "fmt"

func main() {
	// 第一种写法
	//var a int
	//var b bool
	//var c string
	//var d float32

	// 第二种写法
	//var (
	//	a int
	//	b bool
	//	c string
	//	d float32
	//)
	//a = 32
	//b = true
	//c = "string"
	//d = 123.4

	// 第三种写法
	a := 32
	b := true
	c := "string"
	d := 123.4
	fmt.Printf("a-%d b-%t c-%s d-%f\n", a, b, c, d) // 格式化输入

}
