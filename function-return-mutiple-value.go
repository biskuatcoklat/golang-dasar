package main

import "fmt"

func getFullName() (string,string){
	return "Wahyu","Aprilliandhika"
}

func main(){
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}