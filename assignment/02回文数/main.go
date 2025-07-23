package main

import (
	"fmt"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func hwNumber(num int) bool {
	// 方法1：转换成字符串后反转，对原字符串和反转后字符串比较，相同是回文数，不同不是回文数
	// runestr := strconv.Itoa(num)
	// runes := []rune(runestr)
	// origin := runestr
	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 	runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes) == origin

	// 方法2：将数字的前一半和后一半比较，相同是回文数，不同不是回文数
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	revertedNumber := 0
	for num > revertedNumber {
		revertedNumber = revertedNumber*10 + num%10
		num = num / 10
	}
	return revertedNumber == num || num == revertedNumber/10

}

func main() {
	num := 98789
	ok := hwNumber(num)
	if ok {
		fmt.Println(num, "是回文数")
	} else {
		fmt.Println(num, "不是回文数")
	}
}
