package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func OpenCsvFile(fileName string) *csv.Reader {
	// Load a csv file.
	f, _ := os.Open(fileName)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','

	return r
}

func DataToConsole(csvData *csv.Reader) {
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
			fmt.Printf("  %v,", record[value])
		}
		//fmt.Println(record[2]) // display position 2 items
	}
}

// CSVFileToMap  reads csv file into slice of map
// slice is the line number
// map[string]string where key is column name
func CSVFileToMap(csvData *csv.Reader) (returnMap []map[string]string, err error) {

	rawCSVdata, err := csvData.ReadAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	header := []string{} // holds first row (header)
	for lineNum, record := range rawCSVdata {

		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string]string k=header v=value
			line := map[string]string{}
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}
			returnMap = append(returnMap, line)
		}
	}

	return
}

func PrintMap(dataMap []map[string]string) {
	// loop over elements of slice
	for _, m := range dataMap {

		// m is a map[string]interface.
		// loop over keys and values in the map.
		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}
}

func main() {
	f := OpenCsvFile("./testdata_short.csv")
	//DataToConsole(f)
	m, _ := CSVFileToMap(f)

	PrintMap(m)

}
