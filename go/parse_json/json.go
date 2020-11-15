package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("example.json")
	if err != nil {
		fmt.Println("Error opening example.json", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
