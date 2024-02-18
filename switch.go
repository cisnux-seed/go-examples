package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	print("Enter a number: ")
	value, _ := reader.ReadString('\n')

	switch number, _ := strconv.Atoi(value[:len(value)-1]); number%2 != 0 {
	case true:
		fmt.Println("odd number")
	case false:
		println("even number")
	}

	print("Enter a number: ")
	value, _ = reader.ReadString('\n')

	switch number, _ := strconv.Atoi(value[:len(value)-1]); {
	case number%2 == 0:
		println("even number")
	case number%2 != 0:
		fmt.Println("odd number")
	}

}
