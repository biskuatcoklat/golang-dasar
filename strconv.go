package main

import (
	"fmt"
	"strconv"
)

func main() {
	number, err := strconv.ParseInt("10000",10,64)
	if err == nil{
		fmt.Println(number)
	}else{
		fmt.Println(err.Error())
	}

}