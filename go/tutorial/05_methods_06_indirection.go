package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v)

	ScaleFunc(&v, 5)
	fmt.Println(v)

	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p)

	ScaleFunc(p, 8)
	fmt.Println(v, p)
}
