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
	fmt.Println(text) // 结果是15 因为在 Go 中，一个汉字占三个字节
	// 最后出现的位置
	test := strings.LastIndex(a, "晚")
	fmt.Println(test)
	// join
	c := strings.Join(b, "")
	fmt.Println(c)
	// 字符串修改
	var d []rune // rune 只有在使用中文的时候才用,rune代表了一个UTF8字符
	d = []rune(a)
	d[0] = '我'
	a = string(d)
	fmt.Println(a)

	// 非中文字符串反转
	e := "helloGo"
	fmt.Println(e)
	fmt.Println(len(e))
	f := []byte(e)
	fmt.Println(f)
	for i := 0; i <= len(f)/2; i++ {
		tmp := f[i]
		f[i] = f[len(f)-i-1]
		f[len(f)-i-1] = tmp

	}
	fmt.Println(string(f))
	// 中文字符串反转
	g := "中文字符串反转"
	h := []rune(g)
	for i := 0; i < len(h)/2; i++ {
		tmp := h[i]
		h[i] = h[len(h)-i-1]
		h[len(h)-i-1] = tmp
	}
	fmt.Println(string(h))
	// 判断一个字符串是否是回文

	j := "abcdefgfedcba"
	k := []byte(j)
	for i:=0;i<len(k);i++{
		tmp:=k[i]
		k[i]=k[len(k)-i-1]
		k[len(k)-i-1]=tmp

	}
	if string(k)==j{
		fmt.Println("这个字符串是回文")
	}else{
		fmt.Println("这个字符串不是回文")
	}
}
