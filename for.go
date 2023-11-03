package main

import (
	"fmt"
)

func main(){
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("perulangan ke: ", counter)
	// 	counter++
	// }

	for counter:=1; counter<=10; counter++{
		fmt.Println("perulangan ke,", counter)
	}

	slice := []string{"Wahyu","Aprilliandhika","sakti","zaki"}
	// for i := 0; i < len(slice); i++{
	// 	fmt.Println(slice[i])
	// }

	for i, value:= range slice{
		fmt.Println("index",i,"=",value)
	}

	person := make(map[string]string)
	person["nama"]="wahyu"
	person["umur"]="20"

	for key, value:= range person{
		fmt.Println(key,"=",value)
	}
}