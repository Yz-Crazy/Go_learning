package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 新建文件
func Create_File() {
	f, err := os.Create("test.txt") // 新建文件,会创建到根目录

	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	defer f.Close() // 关闭文件

	fmt.Println("文件创建成功")
	// 写入数据
	f.WriteString("你好，这是字符串类型写入\n")
	// byte类型写入，一定要用单引号，而且不能有中文
	b := []byte{'b', 'y', 't', 'e', ' ', 'a', 'b'}
	f.Write(b)
	// slice 写入
	str := "这个是切片类型写入"
	b1 := []byte(str)
	f.Write(b1)
}
func Open_File() {
	f, err := os.Open("test.txt") // 打开文件
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer f.Close()

	// 写入字符串类型数据
	f.WriteString("这是一个啥东西？？？")

	// 文件操作
	buf := make([]byte, 1024)

	// 读取文件内容
	f.Read(buf)
	fmt.Println(string(buf))

}

func Handle_Read_File() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// 这里二选一都可以
		line := scanner.Text()
		// line:=scanner.Bytes()
		fmt.Printf("%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("读取文件报错：%s，err：[%v]", "test.txt", err)
		return
	}

}

func main() {
	//Create_File() // 创建文件
	//Open_File() // 打开文件
	Handle_Read_File()
}
