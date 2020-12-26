package go_type_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlices(t *testing.T) {
	t.Run("Manupulate slices", func(t *testing.T) {
		src := []int{1, 2, 3, 4}
		fmt.Println(src, len(src), cap(src))
		if len(src) != 4 {
			t.Fatal("len(src) should be 4 but ", len(src))
		}
		if cap(src) != 4 {
			t.Fatal("cap(src) should be 4 but ", cap(src))
		}
	})

	// Slice の作成
	// make(型, length, cap)
	t.Run("Initialize slices", func(t *testing.T) {
		src := make([]int, 2, 3)
		fmt.Println(src, len(src), cap(src))
		if len(src) != 2 {
			t.Fatal("len(src) should be 2 but ", len(src))
		}
		if cap(src) != 3 {
			t.Fatal("cap(src) should be 3 but ", cap(src))
		}
	})

	// Slice のコピー
	// 元の slice と同じ長さの slice を作って、 copy
	t.Run("Copy slices", func(t *testing.T) {
		src := []int{1, 2, 3, 4}

		buf := make([]int, len(src))
		copy(buf, src)

		fmt.Println(buf, len(buf), cap(buf))

		if len(src) != 4 {
			t.Fatal("len(buf) should be 4 but ", len(buf))
		}
		if cap(src) != 4 {
			t.Fatal("cap(buf) should be 4 but ", cap(buf))
		}
	})

	// Slice の結合
	// append(slice1, slice2...) と slice2 を展開して結合する
	t.Run("Concat slices", func(t *testing.T) {
		src1, src2 := []int{1, 2, 3}, []int{5, 6}
		concated := append(src1, src2...)
		fmt.Println(concated)
		if !reflect.DeepEqual(concated, []int{1, 2, 3, 5, 6}) {
			t.Fatalf("expected %v but got %v\n", []int{1, 2, 3, 5, 6}, concated)
		}
	})

	// Slice の要素の削除
	// append(slice[:x], slice[x+1:]...) で、x 番目の要素を削除できる
	t.Run("Delete an element from slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}

		// src[:2] -> 1, 2 (close end)
		// src[3:]... -> 4, 5 (open start)
		buf := append(src[:2], src[3:]...)
		fmt.Println(buf)
		if !reflect.DeepEqual(buf, []int{1, 2, 4, 5}) {
			t.Fatalf("expected %v but got %v\n", []int{1, 2, 4, 5}, buf)
		}
	})

	t.Run("Reverse slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5, 6}
		for i := len(src)/2 - 1; i >= 0; i-- {
			opp := len(src) - 1 - i
			src[i], src[opp] = src[opp], src[i]
		}
		fmt.Println(src)
	})

	t.Run("Filter slice", func(t *testing.T) {
		filter := func(n int) bool {
			return n%2 == 0
		}
		src := []int{1, 2, 3, 4, 5}
		dst := src[:0]
		for _, v := range src {
			if filter(v) {
				dst = append(dst, v)
			}
		}
		for i := len(dst); i < len(src); i++ {
			src[i] = 0
		}

		fmt.Println(dst)
		// ゼロ値をいれることで、 src をガベージコレクションの対象にできる
		if !reflect.DeepEqual(dst, []int{2, 4}) {
			t.Fatal("expected ", []int{2, 4}, " but got ", dst)
		}
	})

	t.Run("Split slice", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5, 6, 7}

		size := 2
		dst := make([][]int, 0, (len(src)+size-1)/size)

		for size < len(src) {
			src, dst = src[size:], append(dst, src[0:size:size])
		}
		dst = append(dst, src)
		fmt.Println(dst)
	})
}
