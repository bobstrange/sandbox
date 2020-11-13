package hello_module

import "testing"

func TestHello(t *testing.T) {
	want := "こんにちは世界。"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
