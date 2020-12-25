package main

import "fmt"

type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type HexEmbedded {
	Hex
}

func main() {
	h := Hex(10)
	fmt.Println(h.String())

	h2 := HexEmbedded(20)
	fmt.Println(h2.String())
}
