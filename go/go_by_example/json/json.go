package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	// Tag 付けして JSON のキー名のマッピング
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolJson, _ := json.Marshal(true)
	fmt.Println(string(boolJson))

	intJson, _ := json.Marshal(100)
	fmt.Println(string(intJson))

	floatJson, _ := json.Marshal(3.14)
	fmt.Println(string(floatJson))

	strJson, _ := json.Marshal("gopher")
	fmt.Println(string(strJson))

	sliceJson, _ := json.Marshal([]string{"apple", "banana", "orange"})
	fmt.Println(string(sliceJson))

	mapJson, _ := json.Marshal(map[string]int{"apple": 5, "banana": 7})
	fmt.Println(string(mapJson))

	response1Json, _ := json.Marshal(&Response1{Page: 1, Fruits: []string{"apple", "banana", "orange"}})
	fmt.Println(string(response1Json))
	// {"Page":1,"Fruits":["apple","banana","orange"]}

	response2Json, _ := json.Marshal(&Response2{Page: 1, Fruits: []string{"apple", "banana", "orange"}})
	fmt.Println(string(response2Json))
	// JSON のキー名のマッピング
	// {"page":1,"fruits":["apple","banana","orange"]}

	// JSON data のデコード
	byteData := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 何でも入る map を用意
	var data map[string]interface{}

	if err := json.Unmarshal(byteData, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	// interface の型アサーションで型を確定させる
	num := data["num"].(float64)
	fmt.Println(num)

	// 空インターフェースで型アサーション
	strs := data["strs"].([]interface{})
	// 内部で再度型アサーション
	str1 := strs[0].(string)
	fmt.Println(str1)

	// json についての詳細は
	// https://blog.golang.org/json
}
