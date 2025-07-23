package main

import (
	"fmt"
)

// 切片
func testSlice() {
	arr := [6]int{1, 3, 5, 7, 9}
	arr1 := make([]int, 2, 4)
	fmt.Println(arr1)
	// 前包后不包
	fmt.Println(arr[1:2])
	// 访问长度和容量
	fmt.Println(len(arr), cap(arr))
	// 添加
	arr2 := []int{}
	a2 := append(arr2, 10, 11)
	fmt.Println(a2)
	// 指定位置添加
	arr3 := []int{1, 2, 4, 5}
	arr3 = append(arr3[:2], append([]int{3}, arr3[2:]...)...)
	fmt.Println(arr3)
	// 指定位置删除
	arr4 := []int{1, 2, 6, 3, 4}
	arr4 = append(arr4[:2], arr4[3:]...)
	fmt.Println(arr4)
	// 切片复制
	src1 := []int{1, 2, 3, 4}
	target1 := make([]int, 4, 5)
	copy(target1, src1)
	fmt.Println(target1)
	// 切片触发扩容前，切片一直共用相同的数组
	// 切片触发扩容后，会创建新的数组。并复制这些数据
	// 切片本身是一个特殊的指针，go针对切片类型添加了一些语法糖，方便使用

}

func main() {
	testSlice()
}
