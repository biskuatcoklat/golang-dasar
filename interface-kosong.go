package main

import "fmt"

func ups(i int ) interface{}{
	if i == 1{
		return 1
	}else if i == 2{
		return true
	}else{
		return "oh"
	}
}
func main(){
	var data interface{} = ups(1)
	fmt.Println(data)
}