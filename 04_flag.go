package main

import (
	"flag"
	"fmt"
)

func main() {
	// contoh mengirimkan flag
	// go run 04_flag.go --host=127.0.0.1 --username=admin --password=secret

	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
