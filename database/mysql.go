package database

import "fmt"

var connection string

func init() {
	connection = "PostgresSQL"
	fmt.Println("Database:", connection)
}

func GetDatabase() string {
	return connection
}
