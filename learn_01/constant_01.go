package main

import "fmt"

func main() {

	// 常量基础学习
	// 在常量中，类型是可选的，可写可不写
	// 常量的变量名一般以全大写表示

	// 没有忽略类型
	//const a int = 01234567
	//const b bool = true
	//const c string = "string"
	//const d float32 = 123.456

	// 忽略类型
	//const a = 98765
	//const b = true
	//const c = "string"
	//const d = 3234.222

	const (
		a = 12345
		b = true
		c = "string"
		d = 3234.234
	)

	fmt.Printf("a-%d b-%t c-%s d-%f\n", a, b, c, d) // 格式化输入

	const (
		e = iota // iota 初始默认为0 之后的定于自动+1，iota只在 const 定义中生效
		f
		g
		h
		i = "strings"
		j
		k
		l = iota
		m
		n
		o
		p
		q = 1 << iota // 这里就是把1向左移12位，因为这里的 iota 是自加1
		r
		s
		t
	)
	// 在 Go 语言中 << 左移 和 >> 右移 是在二进制上操作
	// 如 1<<3 则是 把1向左移3位 00000001 向左移3位后为 00000100 转为十进制为4
	// 如 3<<1 则是 把3向左移1位 00000011 向左移1位后为 00000110 转为十进制为6
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

	// 常量练习题

	const (
		A = iota
		B
		C
		D = 8
		E
		F = iota
		G
	)
	const (
		A1 = iota
		A2
	)
	// const 之间的 iota 互不干扰，iota只是在一个const中递增的，重新写一个const之后iota重新计算
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println("A1,A2")
	fmt.Println(A1)
	fmt.Println(A2)
	const ints int = 1234567890
}
