package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // Needs typecast
	var z uint = uint(f) // Needs typecast
	fmt.Println(x, y, z)
}
