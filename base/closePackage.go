package main

import "fmt"

// 闭包

func main() {

	// 声明闭包
	var testclosepackage func(int) int = func(i int) int {
		return i + 1
	}

	fmt.Println(testclosepackage(99))
	// 声明闭包并立即执行
	returnFunc := func() func(int, string) (int, string) {
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()
	r1, r2 := returnFunc(1, "A")
	fmt.Println(r1)
	fmt.Println(r2)

}
