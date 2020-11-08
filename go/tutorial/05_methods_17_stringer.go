package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%v (%v years)", u.Name, u.Age)
}

func main() {
	a := User{"John Doe", 20}
	b := User{"Jane Doe", 20}
	fmt.Println(a, b)
}
