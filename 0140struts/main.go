package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	return &p
}

func main() {
	fmt.Println(person{"zhang", 18, ""})

	// Omitted fields will be zero-valued.
	p := person{name: "Blob", age: 22}
	fmt.Println(p)
	if p.address == "" {
		fmt.Println("p.address is initialized to \"\"")
	}

	fmt.Println(&person{name: "zhang.zzf"})
	fmt.Println(newPerson("feng.zzf"))

	// personPtr is a Point that point to a person
	personPtr := newPerson("zhang.zzf")
	// go will auto dereference the pointer
	fmt.Println(personPtr.name)
	// struct is mutable.
	personPtr.name = "another_name"
	fmt.Println(personPtr.name)
}
