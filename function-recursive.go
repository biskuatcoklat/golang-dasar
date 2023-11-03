package main

import "fmt"

func recursiveFactorial(value int) int {
	if value == 1{
		return 1
	}else{
		return value * recursiveFactorial(value -1)
	}
}
func main(){
	fmt.Println(recursiveFactorial(10))
}