package main

import (
	"fmt"
)

// byte
func byteAndString() {
	var a string = "hello,world"
	var b []byte = []byte(a)
	var c string = string(b)
	fmt.Println("str:", a, ";byte:", b, "c:", c)
}

// rune
func runeType() {
	var a rune = '我'
	fmt.Println(a)
	var b string = "我爱祖国，123"
	var c []rune = []rune(b)
	fmt.Println(c)
	fmt.Println(len(c))
}

// string
func stringType() {
	var s1 string = "我\n爱\n你"
	var s2 string = `我
	爱
	你`
	fmt.Println(s1 == s2)

}

func main() {
	byteAndString()
	runeType()
	stringType()
}
