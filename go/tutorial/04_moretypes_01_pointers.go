package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i         // point to i
	fmt.Println(*p) // read p through the pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i (should be 21)

	p = &j         // point to j
	*p = *p / 20   // divide j through the pointer
	fmt.Println(j) // see the new value of j (should be 1)
}
