package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // length = 6 capacity = 6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // length = 0 capacity = 6 []
	// Extend its length.
	s = s[:4]
	printSlice(s) // length = 4 capacity = 6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // length = 2 capacity = 4 [5 7]

	s = s[:4]
	printSlice(s) // length = 4 capacity = 4 [5 7 11 13]

	s = s[:3]
	printSlice(s) // length = 3 capacity = 4 [5 7 11]
}

func printSlice(s []int) {
	fmt.Printf("length = %d capacity = %d %v\n", len(s), cap(s), s)
}
