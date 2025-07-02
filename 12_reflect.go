package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func main() {
	person := Person{"Joko"}
	personType := reflect.TypeOf(person)
	structField := personType.Field(0)

	fmt.Println(structField.Name)
}
