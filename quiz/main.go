package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type quiz struct {
	question string
	answer   string
}

func run() error {
	filename := flag.String("csv", "quiz.csv", "csv filename containing quiz in format 'question,answer'")
	flag.Parse()
	quizzes, err := openCSV(*filename)
	if err != nil {
		return fmt.Errorf("error parsing quiz: %w", err)
	}
	correct := 0
	for i, q := range quizzes {
		fmt.Printf("Question #%d: %s = ", i+1, q.question)
		var answer string
		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			return fmt.Errorf("error reading answer: %w", err)
		}
		if q.answer == answer {
			correct++
		}
	}
	fmt.Printf("You got %d out of %d\n", correct, len(quizzes))
	return nil
}

func openCSV(filename string) ([]quiz, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %w", err)
	}
	quizzes := make([]quiz, 0)
	for _, record := range records {
		q := quiz{
			question: record[0],
			answer:   record[1],
		}
		quizzes = append(quizzes, q)
	}
	return quizzes, nil
}
