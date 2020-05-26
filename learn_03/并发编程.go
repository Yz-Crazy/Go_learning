package main

import (
	"fmt"
	"runtime"
)

// chan定义方式
// var 变量 chan =make(chan 数据类型)
//单向 chan 定义方式 函数传递的安全考虑
// 接收
// ch := make(chan<- 数据类型)
// 发送
// ch := make(<-chan 数据类型)
// 指定收发的大小
// ch := make(chan 数据类型 大小)

func recive(c chan int) {
	for {

		fmt.Println(<-c)
	}
}

func main() {
	ch := make(chan int)
	//value := <-ch
	go recive(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		//time.Sleep(time.Millisecond)
	}
	fmt.Println(runtime.GOMAXPROCS(0))   // 查看CPU核心
}
