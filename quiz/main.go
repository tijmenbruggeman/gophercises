package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	quizFilePath = flag.String("file", "./assets/problems.csv", "Path to file containing quiz questions")
)

func main() {
	questions, err := loadQuiz(*quizFilePath);
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	for i, line := range questions {
		fmt.Println(i)
		fmt.Println(line.question)
	}

}

type question struct {
	question string
	answer   string
}

func loadQuiz(filePath string) (quiz []question, err error) {
	var questions []question
	file, err := os.Open(filePath)
	if (err != nil) {
		return questions, err
	}
	reader := csv.NewReader(file);
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		question := question{
			question: line[0],
			answer: line[1],
		}
		questions = append(questions, question)
	}

	return questions, err
}
