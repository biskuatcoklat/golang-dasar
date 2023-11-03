package main

import "fmt"

func main(){
	var names[3] string
	names[0] = "Wahyu"
	names[1] = "April"
	names[2] = "Andhika"
	fmt.Println(names[0],names[1])

	var angka = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(angka)
	fmt.Println(angka[1])
	fmt.Println(len(names))
}