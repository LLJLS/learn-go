package main

import (
	"fmt"
)

// if-else if-else

// case
// 不用break

func testCase() {
	a := "test"
	// 可以声明变量
	switch b := 2; a {
	case "t":
		fmt.Println("t")
	case "e":
		fmt.Println("e")
	// 可以匹配多个值
	case "test", "s":
		fmt.Println("test", b)
	default:
		fmt.Println("none")
	}

	// 可以没有条件
	switch {
	case 1 == 1:
		fmt.Println(true)
	case 1 != 1:
		fmt.Println(false)
	default:
		fmt.Println("error")
	}
	// 判定类型
	var c interface{}
	c = 1
	switch d := c.(type) {
	case byte:
		fmt.Println("byte", d)
	case int:
		fmt.Println("int", d)
	default:
		fmt.Println("error", d)
	}
}

func main() {
	testCase()
}
