package main

import "fmt"

func getHello(name string) string{
	if name == "" {
		return "Halo Bro"
	}else{
		return "hello " + name
	}
}

func main(){
	result := getHello("wahyu")
	fmt.Println(result)
	fmt.Println(getHello(""))
}