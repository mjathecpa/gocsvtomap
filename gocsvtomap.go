package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func openCsvFile(fileName string) *csv.Reader {
	// Load a csv file.
	f, _ := os.Open(fileName)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','

	return r*
}

func spitout(csvData *csv.Reader) {
	for {
		record, err := csvData.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
		//fmt.Println(record[2]) // display position 2 items
	}
}

func main() {
	f := openCsvFile("./testdata.csv")
	spitout(f)
}
