package main

import (
	"fmt"
	"time"
)

// select 如果没有case会阻塞
// select 监测 chan 数据流向
// case 必须为IO 操作
// select 对应异步时间处理
// select 超时处理
// 如果多个case都满足条件，会随机选择其中一个来执行

func randomint() chan int {
	ounint := make(chan int)
	go func() {
		i := 0
		for {
			ounint <- i
			time.Sleep(8 * time.Millisecond)
		}
	}()
	return ounint
}

func main() {
	var c1, c2 chan int
	c1, c2 = randomint(), randomint()
	for {
		select {
		case n := <-c1:
			fmt.Println("channel C1 ", n)
		case n := <-c2:
			fmt.Println("channel c2", n)
		case <-time.After(3 * time.Millisecond):
			fmt.Println("获取数据超时")

		}
		time.Sleep(time.Millisecond*1000)
	}
}
