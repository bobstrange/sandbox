package main

type A interface {
	F()
}

type B interface {
	F()
}

func main() {
	var a A
	var b B
	// OK
	a = b

	var f func(A)
	var g func(B)
	// Not ok
	f = g
}
