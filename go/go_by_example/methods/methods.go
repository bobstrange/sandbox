package main

import (
	"fmt"
)

type Rect struct {
	Width, Height int
}

// Area メソッドのレシーバは *Rect
func (r *Rect) Area() int {
	return r.Width * r.Height
}

// レシーバーはポインタでなく値でも良い
func (r Rect) Perim() int {
	return 2*r.Width + 2*r.Height
}

func main() {
	r := Rect{Width: 10, Height: 5}

	fmt.Println("area:", r.Area())
	fmt.Println("perim:", r.Perim())

	// メソッドコールについて、Go は自動的にレシーバーの値・ポインタの変換をしてくれる
	rptr := &r

	fmt.Println("area:", rptr.Area())
	fmt.Println("perim:", rptr.Perim())
}
