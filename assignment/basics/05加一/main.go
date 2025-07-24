package main

import "fmt"

// 问题描述
/*
加一
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。
*/

// 关键
// 对数组从后往前遍历，找到第一个不是9的数字加1，之后的数字全部为0
// 如果数组中全是9，那么要重新创建一个容量加1的数组，并将第一个元素设置为1
// 不能使用将数组转换成数字后+1的方式，因为可能出现溢出
func addOne(arr []int) []int {
	l := len(arr)
	for i := l - 1; i >= 0; i-- {
		if arr[i] != 9 {
			arr[i]++
			for j := i + 1; j < l; j++ {
				arr[j] = 0
			}
			return arr
		}
	}
	// 全为9时
	arr = make([]int, l+1)
	arr[0] = 1
	return arr
}

func main() {
	arr := []int{9, 9, 9}
	fmt.Println(addOne(arr))
}
