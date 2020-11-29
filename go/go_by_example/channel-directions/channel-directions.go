package main

import (
	"fmt"
)

// 変数 pings chan<- string は、チャンネルへの送信しかできない
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 変数 pings は、チャンネルからの受信しかできない
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
