package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now() // 获取当前时间
	fmt.Println(now)
	Year := time.Now().Year() // 年
	//Year := now.Year()
	Month := time.Now().Month() // 月
	//Month := now.Year()
	Day := time.Now().Day()       // 日
	Hour := time.Now().Hour()     // 时
	Minute := time.Now().Minute() // 分
	Second := time.Now().Second() // 秒
	fmt.Println(Year, Month, Day, Hour, Minute, Second)
	// %02d 保留两位，如果不足两位用0填充
	fmt.Printf("%d:%02d:%02d-%02d:%02d:%02d\n", Year, Month, Day, Hour, Minute, Second)

	unix := now.Unix() // 获取当前时间戳,多用于爬虫中
	fmt.Println(unix)

	// 时间戳转当前时间
	timeObjet := time.Unix(unix, 0)
	fmt.Println(timeObjet)

}

// 定时器
func testTicker() {
	ticker := time.Tick(time.Second)
	for i := range ticker { // 每到达ticker的时间就执行一次，可以当心跳包
		fmt.Printf("%v\n", i)
	}
}

// 秒的一些常量
func timeConst() {
	fmt.Printf("纳秒 Nanosecond Duration %d\n", time.Nanosecond)
	fmt.Printf("微秒 Microsecond Micro %d\n", time.Microsecond)
	fmt.Printf("毫秒 Milisecond Mili %d\n", time.Millisecond)
	fmt.Printf("Second %d\n", time.Second)
}

// 格式化(时间的格式化必须写2006年1月2日 15点04分05秒，不然格式化出来的时间是错误的)，格式可以随便，但是时间不能变
func testFormat() {
	now := time.Now()
	//timestr:=now.Format("2006/01/02 15:04:05")
	//timestr:=now.Format("2006/1/2 15:04")
	//timestr:=now.Format("01/2006/02 15:04:05")
	//timestr:=now.Format("01/02/2006 15:04:05")
	//timestr:=now.Format("02/01/2006")
	timestr := now.Format("02-01-2006")
	fmt.Println(timestr)
}

// 练习01，时间格式化
func testFormat2() {
	now := time.Now()
	//方法一：
	timeStr := now.Format("2006/01/02 15:04:05")
	// 方法二：
	time_str := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(timeStr)
	fmt.Println(time_str)
}


// 计算程序运行时间
func testCost() {
	start_time:=time.Now().UnixNano()


	time.Sleep(time.Second*3)
	end_time := time.Now().UnixNano()


	fmt.Println((end_time - start_time)/1000)
}
func main() {

	testTime()
	//testTicker()
	timeConst()
	testFormat()
	testFormat2()
	testCost()
}
