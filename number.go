package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Satu = ", reflect.TypeOf(1))
	fmt.Println("Dua = ", reflect.TypeOf(2))
	fmt.Println("Tiga = ", reflect.TypeOf(3.5))
}
