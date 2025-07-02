package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parse)

	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration: %d\n", duration3)
}
