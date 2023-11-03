package main

import "fmt"

func main(){
	nama := "Wahyu"

	if nama=="Wahyu"{
		fmt.Println("ini benar")
	} else if nama =="zaki"{
		fmt.Println("ini bolehlah")
	}else{
		fmt.Println("error")
	}

	//sort statement
	if length := len(nama); length > 5{
		fmt.Println("nama benar")
	}else{
		fmt.Println("nama panjang")
	}
}