package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// arary や、slice の iteration を range で行える
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"o": "orange", "a": "apple", "b": "banana"}

	// map の iteration も range で行える
	// 順番は不定になる
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// map の key だけでも interation が行える
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 文字列も iteration できる
	// i は開始 byte 数
	// c は rune
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
