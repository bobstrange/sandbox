package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	arr := [3]int{2, 3, 5}
	s = arr[:0]
	printSlice(s)

	s = arr[:2]
	printSlice(s) // length = 2 capacity = 3 [2 3]

	s = append(s, 1)
	printSlice(s)    // length = 3 capacity = 3 [2 3 1]
	fmt.Println(arr) // [2 3 1]

	s = append(s, 7)
	printSlice(s)    // length = 4 capacity = 6 [2 3 1 7]
	fmt.Println(arr) // [2 3 1]
}

func printSlice(s []int) {
	fmt.Printf("length = %d capacity = %d %v\n", len(s), cap(s), s)
}
