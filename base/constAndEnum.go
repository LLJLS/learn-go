package main

import (
	"fmt"
)

// 常量声明
const a = 1

const b int = 1

const (
	c string = "C"
	d bool   = true
)

type AliasInt string

// 枚举
type Sex string

const (
	Male   Sex = "男人"
	FeMale Sex = "女人"
)

// iota
const (
	one int = iota
	two
	three
)

const (
	junary int = iota + 1
	febuary
)

func main() {

	const e = AliasInt("2")
	fmt.Println(e)
	fmt.Println(FeMale)
	fmt.Println(three)
	fmt.Println(febuary)

}
