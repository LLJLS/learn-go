package main

import (
	"fmt"
	"unsafe"
)

// 一级二级n级指针
func npoint() {
	var s1 string = "我是"
	var p *string = &s1
	var p1 **string = &p
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(*p)
	fmt.Println(**p1)
	fmt.Println("*********************************************************")
	var v string = "s"
	var p01 *string = &v
	var p02 **string = &p01
	var p03 ***string = &p02
	fmt.Println(**p02)
	p01 = &s1
	fmt.Println(**p02)
	fmt.Println(***p03)
}

// 指针—>unsafe.Pointer—>uintptr
func uTou() {
	var s1 string = "Abc"
	var p1 *string = &s1
	var c1 unsafe.Pointer = unsafe.Pointer(p1)
	u1 := uintptr(c1)
	u1 += 100
	u2 := (*uint8)(unsafe.Pointer(u1))
	fmt.Println(*u2)

}

func main() {
	npoint()
	uTou()
}
