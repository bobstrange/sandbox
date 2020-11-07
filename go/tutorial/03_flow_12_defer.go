package main

import "fmt"

func main() {
	defer fmt.Println("!!!")
	defer fmt.Println("world")
	fmt.Println("Hello")
}
