package main

import "fmt"

func main() {
	address1 := new(Address)
	address2 := address1

	address2.City = "Jambi"
	address2.Province = "Jambi"
	address2.Country = "Indonesia"
	address2.PostalCode = 36134

	fmt.Println(address1)
	fmt.Println(address2)
}

type Address struct {
	City       string
	Province   string
	Country    string
	PostalCode uint32
}
