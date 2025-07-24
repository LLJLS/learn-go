package main

//136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
//找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singNumber(arr []int) int {
	// 方法一，异或
	// 利用异或相加时，二进制数相同为0,不同为1
	// 0 和 任何数异或都是原数
	// res := arr[0]
	// for i := 1; i < len(arr); i++ {
	// 	res ^= arr[i]
	// }
	// return res

	// res := 0
	// for _, v := range arr {
	// 	res ^= v
	// }
	// return res

	// 方法二，map
	// 遍历数组存入map,在map中有相同值的删除，不同的添加
	m := make(map[int]int)
	for _, v := range arr {
		_, exist := m[v]
		if exist {
			delete(m, v)
		} else {
			m[v] = v
		}
	}
	for k, _ := range m {
		return k
	}
	return 0

}

func main() {

	arr := []int{1, 2, 3, 3, 2, 1, 9}
	println(singNumber(arr))

}
