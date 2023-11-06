package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == ""{
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Wahyu"}
	var sampletype reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampletype.NumField())
	fmt.Println(sampletype.Field(0).Name)
	fmt.Println(sampletype.Field(0).Tag.Get("required"))
	fmt.Println(sampletype.Field(0).Tag.Get("max"))

	fmt.Println(isValid(sample))

}