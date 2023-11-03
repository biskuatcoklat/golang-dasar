package main

import "fmt"

func main(){
	nama := "Wahyu"
	switch nama {
	case "Wahyu":
		fmt.Println("ini benar")
	case "zaki":
		fmt.Println("ini adek")
	default:
		fmt.Println("okeh")
	}

	total := len(nama)
	switch {
	case total >= 5:
		fmt.Println("halo")
	default:
		fmt.Println("false")
	}
}