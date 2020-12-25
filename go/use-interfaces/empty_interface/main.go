package main

import "fmt"

func main() {
	// メソッドセットが空のインターフェース
	var v interface{}

	// 型 int は、interface{} を実装している
	v = 100
	fmt.Println(v)

	// 型 string は、interface{} を実装している
	v = "hoge"
	fmt.Println(v)
}
