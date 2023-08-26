package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
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
			fmt.Printf("  %v,\n", record[value])
		}
		//fmt.Println(record[2]) // display position 2 items
	}
}

func IterateCSV(csvData *csv.Reader){
	for {
		record, err := csvData.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// ... Display all individual elements of the slice.
		for value := range record {
			fmt.Printf("%v\n", record[value])
		}
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

func PrintMap(dataMap []map[string]interface{}) {
	// loop over elements of slice
	for _, m := range dataMap {

		// m is a map[string]interface.
		// loop over keys and values in the map.
		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}
}

func PrintSlice(dataMap []map[string]interface{}) {
	fmt.Println(dataMap[0])
}

func PrintPiece(dataMap []map[string]interface{}) {
	fmt.Println(dataMap[0]["CreditAmount"])
}

func MapFromJson() {
	jsonStr := `{
        "fruits" : {
            "a": "apple",
            "b": "banana"
        },
        "colors" : {
            "r": "red",
            "g": "green"
        }
    }`

	var x map[string]interface{}

	json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x)
}

func JsonFromFile(fileName string) []byte {
	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteResult, _ := io.ReadAll(jsonFile)
	return byteResult
}

func MapFromJsonFile(jsonFile []byte) {
	jsonStr := jsonFile

	var x map[string]interface{}

	json.Unmarshal([]byte(jsonStr), &x)
	fmt.Println(x)
}

// func JsonSlice(jsonFile []byte) {
// 	jsonStr := jsonFile

// 	var x map[string]interface{}

// 	json.Unmarshal([]byte(jsonStr), &x)
// 	fmt.Println(x)
// }

func main() {
	f := OpenCsvFile("./testdata_short.csv")
	//DataToConsole(f)
	IterateCSV(f)
	//m, _ := CSVFileToMap(f)
	//PrintSlice(m)

	// 2023-08-26 - the below works perfect
	//j := JsonFromFile("./sql_statements_short.json")
	//MapFromJsonFile(j)
}
