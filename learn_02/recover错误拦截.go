package main

import "fmt"

// recover程序报错返回错误信息，后续程序还能继续执行
// 当函数使用了错误拦截，当前函数后续的程序不会执行
// 在函数中，当遇到报错，先执行 recover 然后recover 去 panic 获取信息，之后会把当前函数的内存清理干净，然后返回main函数

func Demo(i int) {
	var arr [10]int

	// 设置错误拦截 defer + 匿名函数 + recover
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 123
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}

func main() {
	Demo(10)
	fmt.Println("程序继续执行。。。")
}
