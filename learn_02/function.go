package main

import "fmt"

// 无参数，无返回值函数
func sayHello() {
	fmt.Println("Hello World")
}

// 有参数，有返回值
func add(a int, b int) int {

	return a + b
}

// 多个同类型参数
func add1(a, b int) int {

	return a + b
}

// 多返回值
func calc(a, b int) (int, int) {
	return a + b, a - b
}

//

func testdefer() {
	defer fmt.Println("性感法师")
	defer fmt.Println("在线教学")
	defer fmt.Println("日薪越亿")
	defer fmt.Println("轻松就业")
}

func GetFn() func() {
	fmt.Println("outside")
	return func() {
		fmt.Println("Inside")
	}
}

func GetFn01() func() {
	fmt.Println("这是最后执行")
	return func() {
		fmt.Println("baidu")
	}
}
func testdefer01() {
	defer GetFn()()
	fmt.Println("Here")
}

func main() {
	sayHello()
	c := add(1, 2)
	fmt.Println(c)
	add1(1, 2)
	sum, sub := calc(2, 1)
	fmt.Printf("sum=%d,sub=%d\n", sum, sub)
	// func(){}() 匿名函数后面加一个括号代表是匿名函数调用
	func() { fmt.Println("匿名函数") }()

	testdefer() // 结果 轻松就业，日薪越亿，在线教学，性感法师
	// defer 练习
	testdefer01() // 最后答案    Outside Here Inside    原因是defer入栈先把 Outside 打印了，然后返回了匿名函数入栈了，根据defer特性，在函数结束前执行，所以就先执行了 Here 最后执行了 defer 打印了匿名函数的 Inside
}
