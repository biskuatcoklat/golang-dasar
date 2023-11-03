package main

import "fmt"

type Customer struct{
	Nama, Alamat string
	Age int
}

func (customer Customer) sayHai(nama string){
	fmt.Println("hello",nama,"my name is",customer.Nama)
}

func main(){
	var wahyu Customer
	wahyu.Nama = "Wahyu"
	wahyu.Alamat = "Medan"
	wahyu.Age = 20

	wahyu.sayHai("sakti")

	fmt.Println(wahyu)
	// fmt.Println(wahyu.Nama)

	zaki := Customer{
		Nama: "Zaki",
		Alamat: "Jakarta",
		Age: 6,
	}
	fmt.Println(zaki)
}