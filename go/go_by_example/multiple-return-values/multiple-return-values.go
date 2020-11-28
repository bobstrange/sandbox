package main

import (
	"fmt"
)

// 戻り値が複数ある場合は、型を (int, int) のように表現する
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
