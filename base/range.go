package main

import (
	"fmt"
)

func testRange() {
	// 字符串迭代
	txt := "我爱祖国"
	for i, v := range txt {
		fmt.Println(i, v)
	}

	a := []rune(txt)
	for k, v := range a {
		fmt.Println(k, v)
	}
	// 数组切片迭代
	arr := [10]string{
		"Close friends are truly treasures.Sometimes they know us better than we know ourselves.Their presence reminds us that we are never realy alone.",
		"I knew if I failed I wouldn't reget that,but I knew the one thing that I might reget was not trying.",
	}

	for i, v := range arr {
		fmt.Println(i, "=", v)
	}
	// channel 迭代(待补充)
}

func main() {
	testRange()
}
