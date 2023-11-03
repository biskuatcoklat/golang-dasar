package main

import "fmt"

type Addrees struct{
	City, Province, Country string
}

func ChaangeCountryToIndonesia(address *Addrees){
	address.Country = "indonesia"
}

func main(){
	var address1 = Addrees{"Medan","Sumut","Indonesia"}
	var address2  = &address1

	address2.City = "Siantar"
	*address2 = Addrees{"jakarta","Sumut","Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 = new(Addrees)
	address3.City = "semarang"
	fmt.Println(address3)

	var alamat = Addrees{
		City: "DIY",
		Province: "jogja",
		Country: "",
	}
	var alamatPointer = &alamat
	ChaangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}