package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants
const (
	defaultFilePath  string = "problems.csv" // flag default
	defaultTotalTime int    = 5              // flag default
	totalQuestions   int    = 5
)

type Quiz struct {
	question, answer string
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

// Validate answer from the user
func isCorrectAnswer(quiz Quiz, ansChannel chan<- bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer := scanner.Text()
		_, err := strconv.Atoi(answer)
		if err == nil {
			ansChannel <- (quiz.answer == answer)
			return
		} else {
			fmt.Printf("'%v' is not a valid number!\n", answer)
		}
	}
	ansChannel <- false
}

// Ask Questions
func askQuestions(questions []Quiz, totalTime int) int {
	currentQuestion := 1
	score := 0
	timer := time.NewTimer(time.Duration(totalTime) * time.Second)

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

			// Routine to check the answer
			ansChannel := make(chan bool)
			go isCorrectAnswer(quiz, ansChannel)

			select {
			case <-timer.C:
				fmt.Printf("Time finished! ")
				return score
			case isCorrect := <-ansChannel:
				if isCorrect {
					score++
				}
			}
			currentQuestion++
		}
	}

	return score
}

// Get CLI Arguments
func getArguments() (string, int) {

	filePath := flag.String("filePath", defaultFilePath, "A valid system filepath for the csv file which contains the Questions. Eg: 'problems.csv'")
	totalTime := flag.Int("totalTime", defaultTotalTime, "A valid integer to indicate the total time (in seconds) for the Quiz. Eg: 10 or '10'")

	flag.Parse()
	return *filePath, *totalTime
}

// Driver function
func main() {
	fmt.Println("*************** Welcome to Goopher Quiz :) ***************")

	filePath, totalTime := getArguments()

	f := readFile(filePath)
	defer f.Close()
	csvData := readCSV(f)
	questions := getQuestions(csvData)
	finalScore := askQuestions(questions, totalTime)
	fmt.Printf("Your Score: %v\n", finalScore)
}
