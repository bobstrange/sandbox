package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a) // [0 0 0 0 0] int のゼロ値で埋められる

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 長さは、ビルトインの len() で
	fmt.Println("len:", len(a))

	// 値を設定して初期化できる
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoDim [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDim)
}
