package main

import "fmt"

type T struct {
	F func() string
}

func (t *T) String() string {
	if t.F == nil {
		panic("not implemented")
	}
	return t.F()
}

func main() {
	var s fmt.Stringer = &T{
		F: func() string {
			return "hello"
		},
	}
	fmt.Println(s)

	s = &T{
		F: func() string {
			return "good bye"
		},
	}
	fmt.Println(s)
}
