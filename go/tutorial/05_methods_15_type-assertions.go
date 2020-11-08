package main

import "fmt"

func main() {
	var i interface{} = "Hi"

	s := i.(string) // i が string 型の値を保持していることを assert している
	fmt.Println(s)  // Hi

	s, ok := i.(string)
	fmt.Println(s, ok) // Hi true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64)
	fmt.Println(f) // panic: interface {} is string, not float64
}
