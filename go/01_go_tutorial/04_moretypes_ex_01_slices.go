package main

import (
	"golang.org/x/tour/pic"
)

func getValue(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}

	for i := range s {
		for j := range s[i] {
			switch {
			case j%15 == 0:
				s[i][j] = 255
			case j%5 == 0:
				s[i][j] = 0
			case j%3 == 0:
				s[i][j] = 255
			default:
				s[i][j] = 50
			}
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
