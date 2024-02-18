package main

func main() {
	n := 8
	println("factorial of", n, "is equal to", factorial(n))
}

func factorial(n int) (result int) {
	if n == 0 || n == 1 {
		result = 1
		return result
	}
	result = n * factorial(n-1)
	return result
}
