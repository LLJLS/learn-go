package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
1.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
// 用锁的时候不要在循环中使用defer
func counter(lock *sync.Mutex, num *int) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*num++
		lock.Unlock()
	}
}

/*
2.题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func counter2(num *atomic.Int32) {
	for i := 0; i < 1000; i++ {
		num.Add(1)
	}
}

func main() {
	// Q1
	res := 0
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		go counter(&lock, &res)
	}
	time.Sleep(time.Second)
	fmt.Println(res)
	// Q2
	// var num atomic.Int32
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go counter2(&num)
	// 	wg.Done()
	// }
	// wg.Wait()
	// time.Sleep(time.Second)
	// fmt.Println(num.Load())
}
