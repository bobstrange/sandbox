package main

import "fmt"

func main() {
	// Slice の作成はビルトインの make() を使う
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// ビルトインの append() で、Slice に要素を加えた新しい Slice を返す
	app := append(s, "d")
	// 元の slice はそのまま
	fmt.Println("s is not mutated:", s)
	fmt.Println("append:", app)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))

	// ビルトインの copy(to, from) で、スライスのコピーができる
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice[low:high] で、部分 Slice を返す
	// (low は含む、high は含まない)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// low を指定しないもしくは high を指定しないこともできる
	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 宣言時に初期化もできる
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Slice についての詳細
	// https://blog.golang.org/slices-intro
}
