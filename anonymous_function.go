package main

func main() {
	var blacklist Blacklist = func(s string) bool {
		return s == "anjing"
	}
	register("fajra", blacklist)
	register("anjing", blacklist)
	register("fajra", func(s string) bool {
		return s == "babi"
	})
	register("babi", func(s string) bool {
		return s == "babi"
	})
}

type Blacklist func(string) bool

func register(name string, isBlocked Blacklist) {
	if isBlocked(name) {
		println("You are blocked", name)
	} else {
		println("Welcome!", name)
	}
}
