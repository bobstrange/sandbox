package main

import (
	"fmt"
)

func main() {
	// 2 つまで値をバッファできるチャンネルを作る
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// ※ バッファのサイズより多くの値を送信しようとすると、fatal error が起きる

	fmt.Println(<- messages)
	fmt.Println(<- messages)
}
