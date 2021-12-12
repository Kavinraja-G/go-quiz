package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Kavinraja-G/go-quiz/helpers"
)

// Constants
const totalQuestions int = 5

// Custom type for the Quiz
type Quiz struct {
	question, answer string
}

// Parse all questions from the CSV file to the custom type
func getQuestions(csvData [][]string) []Quiz {
	// Handle no questions in the file
	if len(csvData) == 0 {
		log.Fatal("No Questions are available! :(")
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

// Validate answers from the user
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

// Ask Questions to the user
func askQuestions(questions []Quiz, totalTime int) int {
	currentQuestion := 1
	score := 0

	// Timer to start the duration of the Quiz
	timer := time.NewTimer(time.Duration(totalTime) * time.Second)

	for _, quiz := range questions {
		if currentQuestion > totalQuestions {
			break
		} else {
			fmt.Printf(
				"Question (%v/%v): What is %v?\n", currentQuestion, totalQuestions, quiz.question,
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

// Driver function
func main() {
	fmt.Println("*************** Welcome to Goopher Quiz :) ***************")

	filePath, totalTime := helpers.GetArguments()

	f := helpers.ReadFile(filePath)
	defer f.Close()
	csvData := helpers.ReadCSV(f)
	questions := getQuestions(csvData)
	finalScore := askQuestions(questions, totalTime)
	fmt.Printf("Your Score: %v\n", finalScore)
}
