package main

import "fmt"

func main() {
	v := Ups()
	fmt.Println(v)
}

// Ups -> any is equal to interface{}
func Ups() any {
	return "Ups"
	//return true
	//return 20
}
