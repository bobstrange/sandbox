package go_type_test

import (
	"fmt"
	"testing"
)

type Speaker interface {
	Speak() string
}

type Bob struct{}

type Jane struct{}

func (b *Bob) Speak() string {
	return "WTH"
}

func (j *Jane) Speak() string {
	return "Hi"
}

type SomeFunc func() string

func (s SomeFunc) Speak() string {
	return s()
}

type Person struct{}

func (p *Person) Greet() string {
	return "Hi!"
}
func (p *Person) Footsteps() string {
	return "Pitapat"
}

type PartyPeople struct {
	Person
}

func (p *PartyPeople) Greet() string {
	return "Sup ?"
}

type Greeter interface {
	Greet() string
}

type Footstepper interface {
	Footsteps() string
}

type GreetFootstepper interface {
	Greeter
	Footstepper
}

func TestInterface(t *testing.T) {
	t.Run("interface basic", func(t *testing.T) {
		var s Speaker = &Bob{}
		if text := s.Speak(); text != "WTH" {
			t.Fatal("expected WTH but got ", s.Speak())
		}
		s = &Jane{}
		if text := s.Speak(); text != "Hi" {
			t.Fatal("expected Hi but got ", s.Speak())
		}
	})

	t.Run("function can implement interface", func(t *testing.T) {
		var s Speaker = SomeFunc(func() string {
			return "Test"
		})
		if text := s.Speak(); text != "Test" {
			t.Fatal("expected Test but got ", s.Speak())
		}
	})

	t.Run("empty interface", func(t *testing.T) {
		stringSlice := []string{
			"John",
			"Jane",
		}
		var i interface{} = stringSlice
		fmt.Println(i)

		// cannot use stringSlice (variable of type []string) as []interface{} value
		// var is []interface{} = stringSlice

		is := make([]interface{}, 0, len(stringSlice))

		for _, s := range stringSlice {
			is = append(is, s)
		}
		// ↓ のアサーション通らないのは何故？
		// if !reflect.DeepEqual(is, stringSlice) {
		// 	t.Fatal("is: ", is, " and stringSlice: ", stringSlice, " should be same")
		// }
	})

	t.Run("embed and interface", func(t *testing.T) {
		var gf GreetFootstepper

		gf = &Person{}
		if text := gf.Greet(); text != "Hi!" {
			t.Fatal("expected Hi! but got ", gf.Greet())
		}
		if text := gf.Footsteps(); text != "Pitapat" {
			t.Fatal("expected Pitapat but got ", gf.Footsteps())
		}

		gf = &PartyPeople{}
		if text := gf.Greet(); text != "Sup ?" {
			t.Fatal("expected Hi! but got ", gf.Greet())
		}
		if text := gf.Footsteps(); text != "Pitapat" {
			t.Fatal("expected Pitapat but got ", gf.Footsteps())
		}
	})
}
