package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Sync
	f("direct")

	// Async
	go f("goroutine")

	// 無名関数でも goroutine を実行できる
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 非同期な goroutine 完了まで待ち受ける
	// より、確実な方法は、 WaitGroup を使用する方法 ref: https://gobyexample.com/waitgroups
	time.Sleep(time.Second)
	fmt.Println("done")
}
