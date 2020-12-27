package go_pointer_test

import (
	"fmt"
	"testing"
)

type T struct {
	Number int
	Text   string
}

type Container struct {
	V T
}

func TestContainerValue(t *testing.T) {
	var c Container
	v := c.V
	v.Number = 1
	if c.V.Number != 0 {
		t.Fatal("c.v.Number should be 0 but got ", c.V.Number)
	}
	c.V.Text = "Hi"
	if c.V.Text != "Hi" {
		t.Fatal("c.V.Text should be Hi but got ", c.V.Text)
	}
}

func TestContainerPointer(t *testing.T) {
	c := map[int]T{0: T{}}
	// cannot assign to struct field c[0].Number in map
	// c[0].Number = 1
	c[0] = T{Number: 1}
	fmt.Println(c[0].Number)
	if c[0].Number != 1 {
		t.Fatal("c[0].Number should be 1 but got ", c[0].Number)
	}
}
