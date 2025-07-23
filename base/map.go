package main

import (
	"fmt"
)

func testMap() {
	// 普通声明,无法添加值
	var commonMap map[string]string
	// commonMap["A"] = "a"
	fmt.Println(commonMap)
	// make声明，可以添加值
	makeMap := make(map[string]string, 10)
	makeMap["B"] = "b"
	fmt.Println(makeMap)
	// 遍历
	makeMap["C"] = "c"
	for k, v := range makeMap {
		fmt.Println(k, v)
	}
	// 删除
	delete(makeMap, "C")
	fmt.Println(makeMap)
	// 更新
	makeMap["B"] = "2"
	fmt.Println(makeMap)
	// 异步,需要加锁(后续补充)
	syncMap := make(map[string]string, 10)

	go func() {
		for {
			syncMap["test"] = "test"
			fmt.Println(syncMap)
		}

	}()

	go func() {
		for {
			test := syncMap["test"]
			fmt.Println(test)
		}
	}()

}

func main() {
	testMap()
}
