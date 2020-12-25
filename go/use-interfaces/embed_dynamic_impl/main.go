package main

import "fmt"

type Func func() string

func (f Func) String() string {
	return f()
}

type stringer fmt.Stringer
type T struct {
	stringer
}

func main() {
	t := &T{
		stringer: Func(func() string {
			return "Hi"
		}),
	}
	fmt.Println(t)
}
