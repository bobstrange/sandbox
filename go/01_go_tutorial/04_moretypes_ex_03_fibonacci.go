package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		tmp := a
		a = b
		b = tmp + b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
