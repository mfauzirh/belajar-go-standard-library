package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "eko,kurniawan,khannedy\n" +
		"budi,sejahtera,sujono"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"budi", "sejahtera", "sujono"})
	writer.Flush()
}
