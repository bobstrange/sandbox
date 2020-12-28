package struct_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type person struct {
	name string
	age  int
}

type doc struct {
	person
	createdAt time.Time
}

type docPointer struct {
	*person
	createdAt time.Time
}

func TestStructCompareWithEmbed(t *testing.T) {
	// ref: https://stackoverflow.com/questions/47134293/compare-structs-except-one-field-golang/47134781

	d1createdAt := time.Now()
	d1 := doc{
		person:    person{"John", 20},
		createdAt: d1createdAt,
	}
	time.Sleep(time.Millisecond)

	d2createdAt := time.Now()
	d2 := doc{
		person:    person{"John", 20},
		createdAt: d2createdAt,
	}

	fmt.Println("d1 == d2", d1 == d2)
	fmt.Println("d1.person == d2.person", d1.person == d2.person)

	d3 := docPointer{
		person:    &person{"John", 20},
		createdAt: time.Now(),
	}
	time.Sleep(time.Millisecond)

	d4 := docPointer{
		person:    &person{"John", 20},
		createdAt: time.Now(),
	}

	// struct が pointer を含む場合は reflect.DeepEqual() で比較
	fmt.Println("d3 == d4", d3 == d4)
	fmt.Println("d3.person == d4.person", d3.person == d4.person)
	fmt.Println("reflect.DeepEqual(d3.person, d4.person)", reflect.DeepEqual(d3.person, d4.person))
}

type user struct {
	name      string
	createdAt time.Time
}

func compareUser(a, b user) bool {
	a.createdAt = b.createdAt
	return reflect.DeepEqual(a, b)
}

func compare(a, b *user) bool {
	// 引数のコピーをしないようにすることもできる
	// user 型の変数を準備
	var aCopy = new(user)
	*aCopy = *a
	aCopy.createdAt = b.createdAt
	return reflect.DeepEqual(aCopy, b)
}

func TestCompareWithFunc(t *testing.T) {
	u1 := user{
		name:      "John",
		createdAt: time.Now(),
	}
	u2 := user{
		name:      "John",
		createdAt: time.Now(),
	}
	fmt.Println("u1 == u2", u1 == u2)
	fmt.Println("compareUser(u1, u2)", compareUser(u1, u2))
	fmt.Println("u1", u1)
	fmt.Println("u2", u2)
}

func TestCompareWithFunc2(t *testing.T) {
	u1 := user{
		name:      "John",
		createdAt: time.Now(),
	}
	u2 := user{
		name:      "John",
		createdAt: time.Now(),
	}
	fmt.Println("u1 == u2", u1 == u2)
	fmt.Println("compare(u1, u2)", compare(&u1, &u2))
	fmt.Println("u1", u1)
	fmt.Println("u2", u2)

}
