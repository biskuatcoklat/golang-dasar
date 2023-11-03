package main

import "fmt"

func sayGoodBye(nama string) string{
	return "good Bye" + nama
}

func main(){
	goodbye := sayGoodBye
	result := goodbye("wahyu")
	fmt.Println(result)
	fmt.Println(sayGoodBye("Sakti"))
}