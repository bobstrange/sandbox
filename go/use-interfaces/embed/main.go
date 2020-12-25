package main

import "fmt"

type Foo struct {
	N int
}

type Bar struct {
	Foo
}

type Number struct {
	N int
}

func (n *Number) Get() int {
	return n.N
}

type Number2 struct {
	Number
	N int
}

type Number3 struct {
	*Number
	M int
}

func main() {
	b := Bar{
		Foo: Foo{
			N: 10,
		},
	}
	fmt.Println(b.N)
	fmt.Println(b.Foo.N)

	num := &Number2{
		N:      100,
		Number: Number{N: 1000},
	}
	fmt.Println("num.N", num.N)
	fmt.Println("num.Number.N", num.Number.N)
	fmt.Println("num.Get", num.Get())

	num2 := &Number2{
		N: 100,
	}
	fmt.Println("num2.N", num2.N)
	fmt.Println("num2.Number.N", num2.Number.N)
	fmt.Println("num2.Get", num2.Get())

	num3 := &Number3{
		Number: &Number{N: 10000},
		M:      1000,
	}
	fmt.Println("num3.N", num3.N)
	fmt.Println("num3.Number.N", num3.Number.N)
	fmt.Println("num3.Get", num3.Get())

}
