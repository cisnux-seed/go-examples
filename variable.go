package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age int

	age = -2147483649

	name = "Fajra Risqulla"
	println(name)

	name = "Cisnux"
	fmt.Println(reflect.TypeOf(name))
	println(age)
	fmt.Println(reflect.TypeOf(age))

	var (
		firstName = "Jack"
		lastName  = "Sparrow"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	const timeMillis = 2000
	const no = 1
	const (
		id   = 1
		data = 2
	)
	fmt.Println(id, data)
}
