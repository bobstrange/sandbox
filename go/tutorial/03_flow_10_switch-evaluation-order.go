package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()

	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far :-(")
	}
}
