package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// map 转 json
	m := make(map[string]interface{})
	m["subject"] = "Go语言开发"
	m["students"] = []string{"张三", "李四", "王五", "赵六"}
	m["price"] = 8890.91
	//slices, err := json.Marshal(m) // 不格式化
	slices, err := json.MarshalIndent(m, "", "    ") // 格式化数据 ，最后一个空字符串一般是4个或者8个空格，相当于tab
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	fmt.Println(string(slices))

	// json 转换成map格式，要对接口断言，不然存储数据容易出错
	// json 转 map 的时候用map或者接口都可以
	// json 的 key 对应 map 的 string value 对应 interface{}
	slice := []byte(`{"price":8890.91,"students":["张三","李四","王五","赵六"],"subject":"Go语言开发"}`)
	// 创建map[string]interface{}
	ms := make(map[string]interface{})
	// var ms interface{}
	error := json.Unmarshal(slice, &ms)
	if error != nil {
		fmt.Println("json to map err", error)
		return
	}
	fmt.Println(ms)
	// map 的 value 为 interface 需要进行类型断言来获取数据内容
	// 在断言中，没有切片类型，因为断言只能进行基础数据类型断言，这里的切片是复合类型，他认为是接口切片，所以用 []interface 来断言
	for key, value := range ms {
		switch v := value.(type) {
		//case interface{}:
		//	fmt.Println("interface 类型",key,value)
		case string:
			fmt.Println("string 类型：", key, v)
		//case []string:
		//	fmt.Println("string切片 类型：", key, v)
		case float64:
			fmt.Println("float64 类型：", key, v)
		case []interface{}:
			fmt.Println("[]interface 类型：", key, v)
			for k, v := range v {
				switch vs := v.(type) {
				case string:
					fmt.Println("string 类型", k, vs)

				}
			}
		default:
			fmt.Println(key, v)
			fmt.Printf("%T\n", v)
		}
	}
}
