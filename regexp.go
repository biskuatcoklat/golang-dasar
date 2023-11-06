package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("w([a-z])u")

	fmt.Println(regex.MatchString("Wahyu"))
	fmt.Println(regex.MatchString("wahyu"))
	fmt.Println(regex.MatchString("waHyu"))

	search := regex.FindAllString("Wahyu wahyu waHyu Ini", -1)
	fmt.Println(search)
}