package main

import "fmt"

func main() {
	address1 := Address{
		City:       "Subang",
		Province:   "Jawa Barat",
		Country:    "Indonesia",
		PostalCode: 36134,
	}
	address2 := &address1
	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// without dereference
	//address2 = &Address{
	//	City:     "Jakarta",
	//	Province: "DKI Jakarta",
	//	Country:  "Indonesia",
	//	PostalCode:   28271,
	//}

	// with dereference
	*address2 = Address{
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Country:    "Indonesia",
		PostalCode: 28271,
	}

	fmt.Println(address1) // without dereference the address1 will not change
	fmt.Println(address2)

	fmt.Printf("memory address of address1 is %p\n", &address1)
	fmt.Printf("address2 has memory address of address1 is %p\n", address2)
	fmt.Printf("memory address of address2 is %p\n", &address2)
}

type Address struct {
	City       string
	Province   string
	Country    string
	PostalCode uint32
}
