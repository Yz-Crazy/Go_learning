package main

import "fmt"

// 函数作为函数的返回值(闭包)

func Demo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// 变量的值始终保存在内存中
	f := Demo()
	fmt.Printf("%T\n", f) // 打印 f 的类型
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	// 当重新实例化一个demo的时候，相当于demo在新的内存中存了一份，f和f1互不干扰
	f1 := Demo()
	fmt.Println(f1())
	// 由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包，否则会造成性能问题，有可能会导致内存泄漏，解决的方法是，在退出函数之前，将不使用的局部变量全部删除
	f = nil
	f1 = nil

}
