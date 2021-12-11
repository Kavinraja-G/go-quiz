package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		data = append(data, Quiz{
			question: strings.TrimSpace(questions[0]),
			answer:   strings.TrimSpace(questions[1]),
		})
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

			for scanner.Scan() {
				answer := scanner.Text()
				_, err := strconv.Atoi(answer)
				if err == nil {
					if quiz.answer == answer {
						score += 1
					}
					break
				} else {
					fmt.Printf("'%v' is not a valid number!\n", answer)
				}
			}
			currentQuestion += 1
		}
	}

	return score
}

// Driver function
func main() {
	fmt.Println("*************** Welcome to Goopher Quiz :) ***************")

	f := readFile(filePath)
	defer f.Close()
	csvData := readCSV(f)
	questions := getQuestions(csvData)
	finalScore := askQuestions(questions)
	fmt.Printf("Your Score: %v\n", finalScore)
}
