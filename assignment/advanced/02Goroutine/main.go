package main

import (
	"fmt"
	"time"
)

/*
1.题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func odd() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println("odd:", i)
		}
	}
}

func even() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even:", i)
		}
	}
}

/*
2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

func do(tasks []func()) {
	for i, v := range tasks {
		go func() {
			start := time.Now()
			v()
			end := time.Since(start)
			fmt.Printf("第%d个任务耗时:%d\n", i, end)
		}()
	}
}
func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("one:", i)
	}
}

func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("two:", i)
	}
}

func main() {
	// Q1
	// go odd()
	// go even()
	// time.Sleep(time.Second)

	// Q2
	arrFunc := make([]func(), 2)
	arrFunc[0] = one
	arrFunc[1] = two
	do(arrFunc)
	time.Sleep(time.Second)

}
