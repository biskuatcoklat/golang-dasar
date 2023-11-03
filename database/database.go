package database

import "fmt"

var connection string

func init() {
	fmt.Println("dipanggil")
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}