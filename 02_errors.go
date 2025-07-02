package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "eko" {
		return NotFoundError
	}

	return nil
}

func main() {
	if err := GetById(""); err != nil {
		if errors.Is(err, ValidationError) { // Menggunakan errors.Is untuk mengecek jenis error
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
