package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "root","database")
	var user *string = flag.String("user", "root","database")
	var password *string = flag.String("password", "root","database")

	flag.Parse()
	fmt.Println("host", *host)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
}