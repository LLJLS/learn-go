package main

import "fmt"

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/

func lcp(s []string) string {
	// 方法1：横向扫描
	// if len(s) == 0 {
	// 	return ""
	// }

	// prefix := s[0]
	// for i := 1; i < len(s); i++ {
	// 	prefix = innerLcp(prefix, s[i])
	// }
	// if len(prefix) == 0 {
	// 	return ""
	// }

	// return prefix
	// 方法2：纵向扫描
	// if len(s) == 0 {
	// 	return ""
	// }

	// first := s[0]
	// for x := 0; x < len(first); x++ {
	// 	for y := 1; y < len(s); y++ {
	// 		if x == len(s[y]) || s[y][x] != s[0][x] {
	// 			first = first[:x]
	// 		}
	// 	}
	// }
	// return first

	// 方法3：分治算法
	// if len(s) == 0 {
	// 	return ""
	// }

	// var lcp1 func(int, int) string
	// lcp1 = func(start, end int) string {
	// 	if start == end {
	// 		return s[start]
	// 	}
	// 	mid := (start + end) / 2
	// 	leftlcp, rightlcp := lcp1(start, mid), lcp1(mid+1, end)
	// 	l := min(len(leftlcp), len(rightlcp))
	// 	for i := 0; i < l; i++ {
	// 		if leftlcp[i] != rightlcp[i] {
	// 			return leftlcp[:i]
	// 		}
	// 	}
	// 	return leftlcp[:l]
	// }

	// return lcp1(0, len(s)-1)
	// 方法4：二分查找法
	if len(s) == 0 {
		return ""
	}

	isCommon := func(length int) bool {
		str0, count := s[0][:length], len(s)
		for i := 1; i < count; i++ {
			if s[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	minLength := len(s[0])
	for _, v := range s {
		if len(v) < minLength {
			minLength = len(v)
		}
	}

	low, hight := 0, minLength
	for low < hight {
		mid := (hight-low+1)/2 + low
		if isCommon(mid) {
			low = mid
		} else {
			hight = mid - 1
		}
	}
	return s[0][:low]

}

func innerLcp(s1, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0
	if index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(i1, i2 int) int {
	if i1 <= i2 {
		return i1
	} else {
		return i2
	}
}

func main() {
	s := []string{"leit", "leig", "left"}
	fmt.Println(lcp(s))
}
