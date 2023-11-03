package main

import "fmt"

func getCompleteName()(firstName string, lastName string){
	firstName = "wahyu"
	lastName = "aprilliandhika"
	return
}

func main(){
	f,l := getCompleteName()
	fmt.Println(f,l)
}