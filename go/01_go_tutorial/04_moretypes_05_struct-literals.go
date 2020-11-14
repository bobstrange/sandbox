package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 10} // X:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
	fmt.Println(v1.X)
	fmt.Println(p.Y)
}
