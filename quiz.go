package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Constants
const (
	filePath       string = "problems.csv"
	totalQuestions int    = 5
)

type Quiz struct {
	question string
	answer   string
}

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

// Get all questions
func getQuestions(csvData [][]string) []Quiz {
	// Handle no questions in the file
	if len(csvData) == 0 {
		panic("No Questions are available! :(")
	}

	var data []Quiz
	for _, questions := range csvData {
		q := Quiz{}
		q.question = questions[0]
		q.answer = questions[1]
		data = append(data, q)
	}
	return data
}

// Ask Questions
func askQuestions(questions []Quiz) int {
	currentQuestion := 1
	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	for _, quiz := range questions {
		if currentQuestion > totalQuestions {
			break
		} else {
			fmt.Printf(
				"Question (%v/%v): What is %v?\n",
				currentQuestion,
				totalQuestions,
				quiz.question,
			)
			scanner.Scan()
			answer := scanner.Text()
			if quiz.answer == answer {
				score += 1
			}
			currentQuestion += 1
		}
	}

	return score
}

// Driver function
func main() {
	fmt.Println("Welcome to Goopher Quiz :)")

	f := readFile(filePath)
	defer f.Close() //Closes file at the end of driver function
	csvData := readCSV(f)
	questions := getQuestions(csvData)
	finalScore := askQuestions(questions)
	fmt.Printf("Your Score: %v\n", finalScore)
}
