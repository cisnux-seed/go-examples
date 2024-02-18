package main

import "fmt"

func main() {
	address1 := Address{
		City:       "Subang",
		Province:   "Jawa Barat",
		Country:    "Indonesia",
		PostalCode: 36134,
	}

	//setCity(address1, "Bandung")
	setCity(&address1, "Bandung")
	fmt.Println(address1)
}

func setCity(address *Address, city string) {
	address.City = city
}

type Address struct {
	City       string
	Province   string
	Country    string
	PostalCode uint32
}
