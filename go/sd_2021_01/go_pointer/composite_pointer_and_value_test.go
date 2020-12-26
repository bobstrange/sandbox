package go_pointer_test

import "testing"

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
