package main

import "fmt"

type Hasname interface{
	getName() string
}

func sayHelo(hasname Hasname){
	fmt.Println("hello",hasname.getName())
}

type Person struct{
	Name string
}

type Animal struct{
	Name string
}

func (person Person) getName() string{
	return person.Name
}

func (animal Animal) getName() string{
	return animal.Name
}

func main(){
	var wahyu Person
	wahyu.Name = "wahyu"
	sayHelo(wahyu)

	Dog := Animal{
		Name: "cuk",
	}
	sayHelo(Dog)
}