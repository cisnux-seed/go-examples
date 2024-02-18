package main

import "fmt"

func main() {
	person := Person{
		Name: "Fajra",
	}
	fmt.Println(person)
	sayHello(&person)
	sayHello(createObjectThatHasName())
}

type Person struct {
	Name string
}

func (person *Person) GetName() string {
	return person.Name
}

func sayHello(hasName HasName) {
	println("Hello", hasName.GetName())
}

type HasName interface {
	GetName() string
}

func createObjectThatHasName() (hasName HasName) {
	// any struct that has method, which has a pointer receiver
	// we need to assign reference of struct to variable
	hasName = &Person{Name: "Cisnux"}
	return
}
