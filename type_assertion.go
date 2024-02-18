package main

import "fmt"

func main() {
	//result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//resultInt := result.(int) // panic
	//fmt.Println(resultInt)
	switch result := random(); value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func random() any {
	return "Ups"
	//return 20
	//return true
}
