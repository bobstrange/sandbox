package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Print("loop", i, "value", z, "\n")
		fmt.Print("loop", i, "prev value", prev, "\n")
		if math.Abs(prev-z) < 0.0001 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println("My Sqrt(2):", Sqrt(2))
	fmt.Println("math.Sqrt(2):", math.Sqrt(2))
}
