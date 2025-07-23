package main

import "fmt"

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/

func lcp(s []string) string {
	// 方法1：横向扫描
	if len(s) == 0 {
		return ""
	}

	prefix := s[0]
	for i := 1; i < len(s); i++ {
		prefix = innerLcp(prefix, s[i])
	}
	if len(prefix) == 0 {
		return ""
	}

	return prefix
	// 方法2：纵向扫描

	// 方法3：分治算法
	// 方法4：二分查找法
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
	s := []string{"leit", "leig", "lft"}
	fmt.Println(lcp(s))
}
