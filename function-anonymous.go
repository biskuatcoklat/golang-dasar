package main

import "fmt"

type Blacklist func(string) bool

func registerUser(nama string, blacklist Blacklist){
	if blacklist(nama) {
		fmt.Println("you are bloccked", nama)
	}else{
		fmt.Println("welcome", nama)
	}
}

func main(){
	blacklist := func(nama string) bool{
		return nama == "admin"
	}
	registerUser("admin",blacklist)
	registerUser("Wahyu",blacklist)
}