package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 协程无参直接运行
	go func() {
		fmt.Println("run gorutine in closure")
	}()

	// 协程有参直接运行
	go func(s string) {
		fmt.Println(s)
	}("gorutine:closure params")
	// 协程调用运行
	go say("in goroutine:world")
	// 普通运行
	say("hello")
}
