package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("benar = ", reflect.TypeOf(true))
	fmt.Println("salah = ", reflect.TypeOf(false))
}
