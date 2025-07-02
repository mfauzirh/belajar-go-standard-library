package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	person := Person{"Joko"}
	personType := reflect.TypeOf(person)
	structField := personType.Field(0)

	fmt.Println(structField.Name)

	fmt.Println(IsValid(Person{""}))
}
