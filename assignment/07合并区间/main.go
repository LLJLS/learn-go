package main

import (
	"fmt"
	"sort"
)

// 问题描述
/*
合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(arrarr [][]int) [][]int {
	sort.Slice(arrarr, func(i, j int) bool {
		return arrarr[i][0] < arrarr[j][0]
	})

	res := [][]int{}
	prev := arrarr[0]
	l := len(arrarr)
	for i := 1; i < l; i++ {
		cur := arrarr[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}
func main() {
	arrarr := [][]int{{1, 2}, {3, 5}, {4, 6}}
	fmt.Println(merge(arrarr))

}
