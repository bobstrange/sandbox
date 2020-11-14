package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) // a length = 5 capacity = 5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b length = 0 capacity = 5 []

	c := b[:2]
	printSlice("c", c) // c length = 2 capacity = 5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d length = 3 capacity = 3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s length = %d capacity %d %v\n", s, len(x), cap(x), x)
}
