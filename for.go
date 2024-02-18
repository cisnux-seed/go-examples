package main

import "fmt"

func main() {
	counter := 1

	// equivalent to while in other languages
	for counter <= 10 {
		println(counter, "interation")
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "iteration")
	}

	names := []string{"Cisnux", "Fajra", "Risqulla"}
	for _, name := range names {
		println("name:", name)
	}

	for i := 0; i < 5; i++ {
		for j := 5; j > i-1; j-- {
			print(" ")
		}
		for k := 0; k <= i; k++ {
			if i == 2 && k == 2 {
				print("   ")
			} else {
				print(" * ")
			}

		}
		println()
	}
	for i := 0; i < 5; i++ {
		for k := 0; k <= i+2; k++ {
			if i == 4 {
				break
			}
			print(" ")
		}
		for j := 4; j > i; j-- {
			if i == 0 && j == 4 || i == 0 && j == 3 || i == 1 && j == 4 {
				print("   ")
			} else {
				print(" * ")
			}
		}
		println()
	}
}
