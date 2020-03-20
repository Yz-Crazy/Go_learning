package main

import "fmt"

// 参数为数组的冒泡排序
func BubbleSort(arr [10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素，满足条件交换数据
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}
	fmt.Println(arr)

}

// 参数为切片的冒泡排序
func BubbleSortSlice(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素，满足条件交换数据
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}

	fmt.Println(arr)
	return arr

}

func main() {
	// var 数组名 [元素个数]数据类型
	var arr [10]int
	arr[0] = 123
	fmt.Println(arr)
	var arrs = [10]int{9, 2, 5, 4, 3, 1, 8, 7, 0, 6}
	BubbleSort(arrs)

	// 切片
	// var 切片名 []数据类型
	var silces []int
	silces = append(silces, 1)
	fmt.Println(silces)
	// make([]数据类型，切片长度，切片容量)
	slices := make([]int, 10, 20)
	//slices = append(slices, 1)
	slices[0] = 1
	fmt.Println(slices)
	fmt.Printf("slice=%d,len=%d,cap=%d", slices, len(slices), cap(slices))
	// 内存开辟空间为容量大小，容量可以进行扩容，如果未超过1024 每次扩容为上一次2倍，如果超过1024，为上一次1.25倍
	slice := []int{9, 2, 5, 4, 3, 1, 8, 7, 0, 6}
	BubbleSortSlice(slice)
	fmt.Println(slice)

	// 切片截取
	slicec := []int{9, 2, 5, 4, 3, 1, 8, 7, 0, 6}
	// slice[起始位置数组下标:结束位置下标+1]
	s := slicec[2:6]
	fmt.Println(s)
	// 切片的所有操作都是在原数组上操作
	s[2] = 88
	fmt.Println(s)
	fmt.Println(slicec)

	// 切片的拷贝
	ss := make([]int, 10, 20)
	copy(ss, slicec)
	ss[1]=123
	fmt.Printf("%p\n",ss)
	fmt.Printf("%p\n",slicec)
	fmt.Println(ss)
	fmt.Println(slicec)
}
