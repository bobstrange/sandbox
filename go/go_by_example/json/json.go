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
	// {"page":1,"fruits":["apple","banana","orange"]}

}
