package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Wahyu Aprilliandhika","Wahyu"))
	fmt.Println(strings.Contains("Wahyu Aprilliandhika","sakti"))

	fmt.Println(strings.Split("Wahyu Aprilliandhika", "  "))
	fmt.Println(strings.ToLower("Wahyu Aprilliandhika"))
	fmt.Println(strings.ToUpper("Wahyu Aprilliandhika"))
	fmt.Println(strings.ToTitle("Wahyu Aprilliandhika"))


	fmt.Println(strings.Trim("       Wahyu Aprilliandhika"      , " "))
}