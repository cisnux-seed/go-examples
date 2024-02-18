package main

import "fmt"

func main() {
	user := map[string]string{
		"name":    "Fajra",
		"address": "Jambi",
	}

	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["address"])

	book := make(map[string]string)
	book["title"] = "Go-Lang"
	book["author"] = "Cisnux"
	book["wrong"] = "Ups"
	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}
