package main

import (
	"fmt"
)

func main() {
	// channel の作成は、 make(chan 型) でおこなう
	messages := make(chan string)

	// channel <- で、チャンネルに値を送る
	go func() { messages <- "ping" }()

	// <- channel で、チャンネルから値を受け取る
	msg := <-messages
	fmt.Println(msg)

	// デフォルトでは、送信と受信は、送信側と受信側の準備ができるまでブロックされる。
	// それによって、同期待ちをしなくても、"ping" メッセージの待受ができる
}
