package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // Untuk mendapatkan arguments dari perintah saat menjalankan program
	fmt.Println(args)

	hostname, err := os.Hostname() // Mendapatkan hostname
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
