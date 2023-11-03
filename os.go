package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil{
		fmt.Println("hostname: ", name)
	}else {
		fmt.Println("error", err.Error())
	}
}