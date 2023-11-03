package main

import "fmt"

func main(){
	person := map[string]string{
		"nama" : "Wahyu",
		"alamat":"medan",
	}

	fmt.Println(len(person))
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
}