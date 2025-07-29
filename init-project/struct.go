// go:build linuxx
// + build linux
// 包声明
package main

//引入包
import (
	"fmt"
)

// 函数
func printInConsole(s string) {
	fmt.Println(s)
}

// 全局变量
var str string = "Hello World"

// 入口函数
func main() {
	var a complex64 = 10.9 + 8.7i
	fmt.Println(real(a))
	fmt.Println(imag(a))
	var b uint8 = 8
	var c byte = 8
	fmt.Println(b == c)
	printInConsole(str)
}
