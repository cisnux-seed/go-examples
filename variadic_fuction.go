package main

func main() {
	println("total: ", sum(10, 20, 30))
	numbers := []int{10, 20, 30, 60}
	println("total: ", sum(numbers...))

}

func sum(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return total
}
