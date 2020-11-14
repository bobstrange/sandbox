package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("10 + 20 =", add(10, 20))
}
