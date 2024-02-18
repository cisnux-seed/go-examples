package main

func main() {
	firstName, lastName := getFullName()
	println(firstName, lastName)
}

func getFullName() (firstName, lastName string) {
	firstName = "Fajra"
	lastName = "Risqulla"
	return
}
