package main

import "fmt"

type Func func() string

func (f Func) String() string {
	return f()
}

func main() {
	var s fmt.Stringer = Func(
		func() string {
			return "Hi"
		},
	)
	fmt.Println(s)
}
