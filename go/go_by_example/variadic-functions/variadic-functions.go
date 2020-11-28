package main

import (
	"fmt"
)

// 可変長の引数を受け取る関数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 変数の後ろに ... をつけて slice を展開することもできる
	sum(nums...)
}
