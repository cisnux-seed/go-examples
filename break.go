package main

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			break
		}
		println(i, "is a odd number")
	}
}
