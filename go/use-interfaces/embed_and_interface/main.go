package main

import "fmt"

type Gender int

const (
	Female = iota
	Male
)

type Person interface {
	Name() string
	Title() string
}

func NewPerson(gender Gender, firstName, lastName string) Person {
	p := &person{firstName: firstName, lastName: lastName}
	if gender == Male {
		return &male{p}
	} else {
		return &female{p}
	}
}

type person struct {
	firstName string
	lastName  string
}

func (p *person) Name() string {
	return p.firstName + " " + p.lastName
}

type male struct {
	*person
}

type female struct {
	*person
}

func (f *female) Title() string {
	return "Ms."
}

func (m *male) Title() string {
	return "Mr."
}

func printName(p Person) {
	fmt.Println(p.Title(), p.Name())
}

func main() {
	p1 := NewPerson(Male, "John", "Doe")
	printName(p1)

	p2 := NewPerson(Female, "Jane", "Doe")
	printName(p2)
}
