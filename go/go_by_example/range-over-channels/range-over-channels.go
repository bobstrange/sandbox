package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range で、channel 内の値を一つずつ取ってくることができる
	// channel が close されていないと all goroutines are asleep error になる
	for elem := range queue {
		fmt.Println(elem)
	}
}
