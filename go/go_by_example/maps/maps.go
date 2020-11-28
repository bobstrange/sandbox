package main

import (
	"fmt"
)

func main() {
	// ビルトインの make で生成
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13
	fmt.Println("map: ", m)

	// len() では、key value のペアの個数を返す
	fmt.Println("len: ", len(m))

	delete(m, "key2")
	fmt.Println("map:", m)

	// 戻り値の 2 番目は、map にキーが存在していたかどうかを返す
	_, prs := m["key2"]
	fmt.Println("prs: ", prs)

	// {} で初期化もできる
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
