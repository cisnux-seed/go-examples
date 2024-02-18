package main

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Fajra", spamFilter)
}

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	println("Hello", filter(name))
}

func spamFilter(name string) (filteredValue string) {
	if name == "Anjing" {
		return "..."
	}
	return name
}
