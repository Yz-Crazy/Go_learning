package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	// 在转换成json的时候，结构体中的成员首字母必须大写
	Subject  string
	Students []string
	Price    float64
}

func main() {
	// 结构体转换成json
	cl := Class{"Go语言开发", []string{"张三", "李四", "王五", "赵六"}, 8890.91}
	slices, err := json.Marshal(cl) //全部打印成一行
	//slices, err := json.MarshalIndent(cl, "", " ") // 结构化
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	fmt.Println(string(slices))

	// json 转换成 结构体
	slice := []byte(`{"Subject":"Go语言开发","Students":["张三","李四","王五","赵六"],"Price":8890.91}`)
	var cls Class
	error := json.Unmarshal(slice, &cls)
	if error != nil {
		fmt.Println("json err", error)
		return
	}
	fmt.Println(cls)
}
