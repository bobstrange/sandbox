package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "ワン！"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "ニャー"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "????"
}

type Gopher struct {
}

func (g Gopher) Speak() string {
	return "Go is good!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, Gopher{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
