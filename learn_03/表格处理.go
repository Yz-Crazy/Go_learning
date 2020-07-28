package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

// 创建表格
func Create_Xlsx() {
	// 新建表格
	file := xlsx.NewFile()
	// 创建页面
	file.AddSheet("test")
	// 保存文件
	err := file.Save("test.xlsx")
	if err != nil {
		fmt.Println("创建表格失败", err.Error())
		return
	}
	fmt.Println("创建表格成功")
}

// 表格操作

type Cell struct {
	Name []string
}
type Row struct {
	Cell
}
type Sheet struct {
	Sname string // sheet名字
	Row          // 行内容
}

func Create_Excel(sheet Sheet) {
	// 新建表格
	file := xlsx.NewFile()
	// // 设置页面名称
	save, err := file.AddSheet(sheet.Sname)
	if err != nil {
		fmt.Println("写入表格页面报错", err.Error())
	}
	// 添加行
	row := save.AddRow()
	// 设置行高，不填用默认
	row.SetHeightCM(1.5)
	// 循环添加数据
	for _, w := range sheet.Name {
		cell := row.AddCell()
		cell.Value = w
	}
	// 添加数据



	// 保存文件
	err=file.Save("ceesla.xlsx")
	if err != nil{
		fmt.Println("保存失败",err.Error())
	}

}






func main() {
	//Create_Xlsx()      // 创建表格

	sheet := Sheet{"Go语言练习", Row{Cell{[]string{"姓名","学历","高校","行业","工作年限","职位","薪资","编程语言"}}}}
	Create_Excel(sheet)

}

