package main

import (
	"fmt"
	"strconv"
)

func changeType() {
	// 数字转换
	a := 257
	b := byte(a)
	fmt.Println(b)

	// 字符串转换
	// 字符串转数字
	str := "256"
	c, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("%T\n", c)
	}

	// 数字转字符串
	d := strconv.Itoa(c)
	fmt.Printf("%T\n", d)

	// 字符串转uint64
	str1 := "5201314"
	e, err := strconv.ParseUint(str1, 10, 32)
	fmt.Printf("%T\n", e)
	// uint64转字符串
	str2 := strconv.FormatUint(e, 2)
	fmt.Printf("%T\n", str2)
	// 接口类型转换
	var i interface{} = 3
	f, ok := i.(int)
	if ok {
		fmt.Printf("%T\n", f)
	} else {
		fmt.Printf("转换失败")
	}
	// 结构体类型转换
	// 转换条件很苛刻，不仅要名字一样，类型一样，顺序也要一样

}

func main() {
	changeType()
}
