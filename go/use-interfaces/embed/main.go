package main

import "fmt"

type Walker interface {
	Walk()
}

type Dog struct {
}

func (d *Dog) Walk() {
	fmt.Println("Dog walking ...")
}

type Cat struct {
}

func (c *Cat) Walk() {
	fmt.Println("Cat walking...")
}

type Kerberos struct {
	Dog
}

func (k *Kerberos) Bless() {
	fmt.Println("Kerberos blessing...")
}

func (k *Kerberos) Walk() {
	k.Bless()
}

func main() {
	var c, d, k Walker
	c = &Cat{}
	c.Walk()

	d = &Dog{}
	d.Walk()
	d2, ok := d.(*Kerberos)
	if ok == true {
		fmt.Println("d is Kerberos")
		d2.Bless()
	}

	k = &Kerberos{}
	k2, ok := k.(*Kerberos)
	if ok == true {
		fmt.Println("k is Kerberos")
		k2.Bless()
	}
}
