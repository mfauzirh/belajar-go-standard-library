package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Fauzi", "Fauzi"))
	fmt.Println(strings.Split("Muhammad Fauzi", " "))
	fmt.Println(strings.ToLower("Muhamamd Fauzi"))
	fmt.Println(strings.ToUpper("Muhamamd Fauzi"))
	fmt.Println(strings.Trim("    Muhamamd Fauzi  ", " "))
	fmt.Println(strings.ReplaceAll("Cut Paw Cat Paw", "Paw", "Damn"))
}
