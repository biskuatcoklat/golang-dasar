package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(nama string, filter Filter){
	filteredName := filter(nama)
	fmt.Println("Hello ",filteredName)
}
// func sayHelloWithFilter(nama string, filter func(string) string){
// 	filteredName := filter(nama)
// 	fmt.Println("Hello ",filteredName)
// }

func spamFilter(nama string) string{
	if nama == "Anjing"{
		return "..."
	}else{
		return nama
	}
}

func main(){
	sayHelloWithFilter("Wahyu",spamFilter)
	sayHelloWithFilter("Anjing",spamFilter)
	filter := spamFilter
	sayHelloWithFilter("wow",filter)
}