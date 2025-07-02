package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Eko", "Paw"}
	values := []int{10, 20, 5, 2}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "paw"))
}
