package main

import "fmt"

func logging(){
	fmt.Println("aplikasi sukses")
}

func RunTimeApplications(value int){
	defer logging()
	fmt.Println("Run Application")
	result := 10/value
	fmt.Println("result",result)
}

func main(){
	RunTimeApplications(10)
}