package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func newPerson(name string) *Person {
	p := Person{Name: name}
	p.Age = 42

	// function scope のローカル変数 p は、ガベージコレクションされないので、参照を返しても安全
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{Name: "Alice", Age: 30})

	fmt.Println(Person{Name: "Fred"}) // 初期化していないフィールドは zero 値で初期化される

	fmt.Println(&Person{Name: "Ann", Age: 40})

	fmt.Println(newPerson("John"))

	s := Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	// Struct のポインタでも、.Field でフィールドにアクセスできる
	sptr := &s
	fmt.Println(sptr.Age)

	// Struct は Mutable
	sptr.Age = 51
	fmt.Println(sptr.Age)
}
