package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names [3]string
	names[0] = "Cisnux"
	names[1] = "Rajra"
	names[2] = "Risqulla"

	for _, name := range names {
		println(name)
	}

	var numbers = [3]uint8{
		90,
		80,
		70,
	}
	//numbers[2] = -1

	for _, value := range numbers {
		println(value)
	}

	fmt.Println(reflect.TypeOf(numbers))

	var numbers2 = [...]uint8{
		90,
		80,
		70,
		20,
	}

	for _, number := range numbers2 {
		println(number)
	}
}
