package main

import (
	"fmt"
	"time"
)

// done チャンネルはコールバックみたいなイメージ
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// channel から受信することで、worker の完了を待ち受ける
	<-done
}
