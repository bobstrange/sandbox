package main

import (
	"fmt"
	"sort"
)

// sort.Sort() に渡す用の型を作る
type byLength []string

// Len, Less, Swap を実装する
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "dragon-fruits","orange", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
