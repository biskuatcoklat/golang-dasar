package main

import "fmt"

func sayHello(){
	fmt.Println("hello")
}

func main(){
	sayHello()
	for i := 0; i < 10; i++{
		sayHello()
	}
}