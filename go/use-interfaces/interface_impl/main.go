package main

import "fmt"

type Func func() string

func (f Func) String() string {
	return f()
}

func main() {
	var text fmt.Stringer = Func(nil)
	fmt.Println(text)
}
