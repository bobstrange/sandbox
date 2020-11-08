package main

import "fmt"

func main() {
	var i interface{} // 空のインターフェース
	describe(i)

	i = 42 // interface{} 型に assign 可能
	describe(i)

	i = "hello" // interface{} 型に assign 可能
	describe(i)

	i = func() int { // interface{} 型に assign 可能
		return 42
	}
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
