package main

import (
	"fmt"
	"reflect"
)

func main() {
	var integer32 int32 = 32768
	var integer64 = int64(integer32)
	var integer16 = int16(integer32)

	println(integer32)
	println(integer64)
	println(integer16)

	var name = "Fajra Risqulla"
	var char byte = name[4]
	var charToString = string(char)

	println(char)
	fmt.Println(reflect.TypeOf(char))
	println(charToString)

	type MyAge byte

	var age MyAge = 255
	println(age)
	fmt.Println(reflect.TypeOf(age))
}
