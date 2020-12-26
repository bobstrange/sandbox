package go_pointer_test

import "testing"

type Value struct {
	content [64]byte
}

// go:noinline
func f(v Value) Value {
	return v
}

// go:noinline
func g(v *Value) *Value {
	return v
}

func BenchmarkValue(b *testing.B) {
	b.ReportAllocs()
	var v Value
	for i := 0; i < b.N; i++ {
		f(v)
	}
}

func BenchmarkPointer(b *testing.B) {
	b.ReportAllocs()
	var v Value
	for i := 0; i < b.N; i++ {
		g(&v)
	}
}

/*
go test ./go_pointer -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/bobstrange/sandbox/go/sd_2021_01/go_pointer
BenchmarkValue-12       1000000000               0.504 ns/op           0 B/op          0 allocs/op
BenchmarkPointer-12     1000000000               0.250 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/bobstrange/sandbox/go/sd_2021_01/go_pointer  0.848s
*/
