package helpers

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

const (
	defaultFilePath  string = "problems.csv" // flag default
	defaultTotalTime int    = 5              // flag default
)

// Get CLI Arguments
func GetArguments() (string, int) {

	filePath := flag.String("filePath", defaultFilePath, "A valid system filepath for the csv file which contains the Questions. Eg: 'problems.csv'")
	totalTime := flag.Int("totalTime", defaultTotalTime, "A valid integer to indicate the total time (in seconds) for the Quiz. Eg: 10 or '10'")

	flag.Parse()
	return *filePath, *totalTime
}

// Read a file from local path
func ReadFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}

	return f
}

// Read a CSV file and return the data
func ReadCSV(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	return csvData
}
