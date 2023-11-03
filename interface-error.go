package main

import (
	"errors"
	"fmt"
)

func Pemabagi(nilai int, pembagi int)(int,error){
	if pembagi == 0 {
		return 0, errors.New("error dilarang bagi 0")
	}else{
		result := nilai / pembagi
		return result, nil
	}
}

func main(){
	hasil, err := Pemabagi(100,20)
	if err == nil{
		fmt.Println("hasil",hasil)
	}else{
		fmt.Println("error",err.Error())
	}
}