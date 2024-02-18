package main

import "fmt"

func main() {
	address1 := Address{}
	address1.SetCity("Subang")
	address1.SetProvince("Jawa Barat")
	address1.SetPostalCode(36134)
	address1.SetCountry("Indonesia")
	fmt.Println(address1)
}

type Address struct {
	city       string
	province   string
	country    string
	postalCode uint32
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) SetCity(city string) {
	a.city = city
}

func (a *Address) Province() string {
	return a.province
}

func (a *Address) SetProvince(province string) {
	a.province = province
}

func (a *Address) Country() string {
	return a.country
}

func (a *Address) SetCountry(country string) {
	a.country = country
}

func (a *Address) PostalCode() uint32 {
	return a.postalCode
}

func (a *Address) SetPostalCode(postalCode uint32) {
	a.postalCode = postalCode
}
