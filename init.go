package main

import (
	"fmt"
	"go-examples/database"
	// blank identifier only import side effects
	_ "go-examples/repository"
)

func main() {
	fmt.Println(database.GetDatabase())
}
