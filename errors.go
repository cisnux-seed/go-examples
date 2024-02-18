package main

import (
	"errors"
	"fmt"
)

func main() {
	x, y := 9, 3
	result, e := division(x, y)
	if e != nil {
		println(e.Error())
	} else {
		fmt.Printf("%d divided by %d is equal to %d\n", x, y, result)
	}
}

func division(x int, y int) (result int, e error) {
	if y == 0 {
		e = errors.New("can't divided by zero")
	} else {
		result = x / y
	}
	return
}
