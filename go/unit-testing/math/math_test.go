package math

import "testing"

type example struct {
	arg1, arg2, want int
}

var examples = []example{
	{2, 3, 5},
	{4, 8, 12},
	{-1, 2, 1},
	{-1, -2, -3},
}

func TestAdd(t *testing.T) {
	for _, test := range examples {
		if output := Add(test.arg1, test.arg2); output != test.want {
			t.Errorf("output %q not equal to want %q", output, test.want)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
