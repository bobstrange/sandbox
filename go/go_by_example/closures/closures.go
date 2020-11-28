package main

import (
	"fmt"
)

// 関数を返す関数
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	// nextInt を実行するたびに束縛されている変数 i の値がインクリメントされる
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// intSeq() によって返される関数の state は独立
	newInts := intSeq()
	fmt.Println(newInts())
}
