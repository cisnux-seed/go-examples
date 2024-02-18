package main

func main() {
	greeting := sayGoodBye
	println(greeting("Fajra"))
}

func sayGoodBye(name string) string {
	return "Good Bye " + name
}
