package main

import "fmt"

/*
两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。
*/

func twoSum(arr []int, target int) []int {
	// 方法1：暴力枚举
	// l := len(arr)
	// for i := 0; i < l-1; i++ {
	// 	for j := i + 1; j < l; j++ {
	// 		if arr[i]+arr[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return nil

	// 方法2：查找表法
	tmpMap := map[int]int{}
	for i, v := range arr {
		if k, ok := tmpMap[target-v]; ok {
			return []int{i, k}
		}
		tmpMap[v] = i
	}
	return nil
}

func main() {
	arr := []int{3, 4, 3, 3, 7}
	target := 9
	fmt.Println(twoSum(arr, target))

}
