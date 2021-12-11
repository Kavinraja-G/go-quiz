package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Constants
const filePath string = "problems.csv"

// Read a file from local path
func readFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

// Read a CSV file and return the data
func readCSV(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return csvData
}

// Driver function
func main() {
	fmt.Println("Welcome to Go Quiz!")

	f := readFile(filePath)
	defer f.Close() //Closes file at the end of driver function
	csvData := readCSV(f)

	fmt.Println(csvData)
}
