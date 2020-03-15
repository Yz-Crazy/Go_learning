package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split 的应用
	a := "现在已经很晚了晚了已经"
	b := strings.Split(a, "")
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	for k, v := range a {
		fmt.Println(k, v)
		fmt.Printf("%[1]c,%[1]d\n", v)
	}
	fmt.Println(b)

	// 判断是否总在于字符串
	if strings.Contains(a, "晚") {
		fmt.Println("存在于字符串中")
	} else {
		fmt.Println("这个字不存在于字符串中")
	}

	// 前缀判断
	if strings.HasPrefix(a, "现在") {
		fmt.Println("存在于最开始")
	} else {
		fmt.Println("这个字不存在于字符串开头")
	}
	// 后缀判断
	if strings.HasSuffix(a, "已经") {
		fmt.Println("存在于结尾")
	} else {
		fmt.Println("这个字不存在于字符串结尾")
	}
	// 子串最先出现位置
	text := strings.Index(a, "晚")
	fmt.Println(text)   // 结果是15 因为在 Go 中，一个汉字占三个字节
	// 最后出现的位置
	test:=strings.LastIndex(a,"晚")
	fmt.Println(test)
	// join
	c:=strings.Join(b,"")
	fmt.Println(c)
}
