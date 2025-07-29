package main

import (
	"fmt"
	"time"
)

/*
1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func send(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	defer close(ch)
}

func receive(res int) {
	fmt.Println(res)

}

/*
2.题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func product(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	defer close(ch)
}

func consumer(res <-chan int) {
	for v := range res {
		fmt.Println(v)
	}
}

// select中的参数需要是数据，而不是通道本身
func main() {
	// Q1
	ch := make(chan int, 11)
	go send(ch)
	// go receive(ch)
	for i := 0; i < 11; i++ {
		select {
		case data := <-ch:
			go receive(data)
			// default:
			// 	fmt.Println("通道中没有数据了")
		}
	}

	time.Sleep(time.Second)

	// Q2
	// ch1 := make(chan int, 10)
	// go product(ch1)
	// go consumer(ch1)
	// time.Sleep(time.Second)

}
