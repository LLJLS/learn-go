package main

import (
	"fmt"
)

// 循环结构
func testLoop() {
	// fori
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// foreach
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	for {
		break
	}
	// 遍历数组
	var arr [10]string
	arr[0] = "Hello"
	for i, e := range arr {
		fmt.Println(i, "=", e)
	}

	for i, e := range arr[0] {
		fmt.Println(i, "=", (string)(e))
	}
	// 遍历切片
	var arr1 []string = make([]string, 10)
	arr1[0] = "World"
	for i, e := range arr1[0] {
		fmt.Println(i, "=", (string)(e))
	}
	// 遍历map
	// 乱序
	m := make(map[string]string, 10)
	m["one"] = "1"
	m["two"] = "2"
	for k, v := range m {
		fmt.Println(k, "=", v)
	}
	// 使用标记，continue 也一样
out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				fmt.Println("退出标记")
				break out
			}
		}
	}
	// goto 语句,不建议使用

}
func main() {
	testLoop()
}
