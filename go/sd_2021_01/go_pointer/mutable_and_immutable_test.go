package go_pointer_test

import (
	"fmt"
	"testing"
)

type ex struct {
	Number int
}

func TestSliceIteration(t *testing.T) {
	s := []ex{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	s2 := []*ex{}
	for _, v := range s {
		s2 = append(s2, &v)
	}
	for _, v := range s2 {
		fmt.Printf("%+v\n", v)
		// 出力結果
		// &{Number:5}
		// &{Number:5}
		// &{Number:5}
		// &{Number:5}
		// &{Number:5}

		// 理由は、iteration で v は使いまわされている為
		// pointer が同じ
	}
}

func TestSliceIteration2(t *testing.T) {
	s := []*ex{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	s2 := []*ex{}
	for _, v := range s {
		s2 = append(s2, v)
	}
	for _, v := range s2 {
		fmt.Printf("%+v\n", v)
		// 出力結果
		// &{Number:1}
		// &{Number:2}
		// &{Number:3}
		// &{Number:4}
		// &{Number:5}
	}
}

func TestSliceIteration3(t *testing.T) {
	s := []ex{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	s2 := []*ex{}
	for _, v := range s {
		v := v
		s2 = append(s2, &v)
	}
	for _, v := range s2 {
		fmt.Printf("%+v\n", v)
		// 出力結果
		// &{Number:1}
		// &{Number:2}
		// &{Number:3}
		// &{Number:4}
		// &{Number:5}

		// 新しく v をループ内で定義している
		// 各 iteration append される v のアドレスが異なる状態になる
	}
}
