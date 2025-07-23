package main

import (
	"fmt"
)

func testArr() {
	// 声明数组
	var arr1 [10]int
	fmt.Println(arr1)
	// 声明并初始化数组
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	// 不定义数组长度
	var arr3 = [...]int{12, 3, 4}
	fmt.Println(len(arr3))
	print := func(arr [3]int) {
		fmt.Println(arr)
	}
	print(arr3)
	// 精准初始化
	arr4 := []string{2: "A", 3: "B"}
	fmt.Println(arr4)

	// 获取数组中的值
	three := arr2[2]
	fmt.Println(three)

	// 遍历
	for i, v := range arr3 {
		fmt.Println(i, "=", v)
	}
	// 多维数组
	arrarr := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(arrarr)

	// 多维遍历
	narr := [3][2][2]int{
		{
			{111, 112}, {121, 122},
		},
		{
			{211, 212}, {221, 222},
		},
		{
			{311, 312}, {321, 322},
		},
	}
	for i, v := range narr {
		fmt.Println(i, "=", v)
		for j, x := range v {
			fmt.Println(j, "=", x)
			for z, c := range x {
				fmt.Println(z, "=", c)
			}
		}
	}

	// 数组作为参数传递给函数，函数内是修改了，但是无法修改数组，除非传入指针
	var arr5 []int = []int{123, 345, 678}
	// testArrParm := func(a []int) {
	// 	a[1] = 999
	// 	fmt.Println(a)
	// }

	// testArrParm(arr5)
	testArrParm2(arr5)
	fmt.Println(arr5)

}

func testArrParm2(a []int) {
	a[1] = 888
	fmt.Println(a)
}

func main() {
	testArr()
}
