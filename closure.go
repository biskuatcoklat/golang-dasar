package main

import "fmt"

func main(){
	nama := "wahyu"
	counter := 0

	inc := func(){
		fmt.Println("inc")
		counter++
	}
	inc()
	inc()
	fmt.Println(nama)
	fmt.Println(counter)
}