package main

import "fmt"

// 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。

func changeInt(i *int) {
	*i += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func changeArr(arrp *[]int) {
	arr := *arrp
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func main() {
	a := 10
	changeInt(&a)
	fmt.Println("乘以10后的值：", a)
	b := []int{2, 3, 4}
	changeArr(&b)
	fmt.Println("数组中每个数乘以2后的数组:", b)

}
