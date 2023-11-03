package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Jomblo(){
	man.Name = "Mr. " + man.Name
}

func main(){
	w := Man{"Wahyu"}
	w.Jomblo()

	fmt.Println(w.Name)
}