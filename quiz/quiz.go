package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Quiz represents a quiz
type Quiz struct {
	question string
	answer   string
}

// Quizzes represents a slice of quiz
type Quizzes []Quiz

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	limit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order")
	flag.Parse()

	csvFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Failed to read csv file.", err)
	}

	quizzes := getQuizzes(csvFile)
	if *shuffle {
		shuffleQuizzes(quizzes)
	}

	timeLimit := time.Duration(*limit) * time.Second
	point := askQuizzes(quizzes, timeLimit, os.Stdout, os.Stdin)

	fmt.Printf("\nCorrect answer(s): %d\n", point)
}

func getAnswer(ch chan string, r io.Reader) {
	answer := ""
	fmt.Fscanf(r, "%s", &answer)
	ch <- strings.TrimSpace(answer)
}

func checkAnswer(answer, solution string, point int) int {
	if answer == solution {
		point++
	}
	return point
}

func quizString(index int, question string) string {
	return fmt.Sprintf("Question #%d: %s = ", index+1, question)
}

func askQuizzes(quizzes Quizzes, timeLimit time.Duration, w io.Writer, r io.Reader) int {
	point := 0
	ch := make(chan string)
	for index, quiz := range quizzes {
		fmt.Fprintf(w, quizString(index, quiz.question))

		go getAnswer(ch, r)

		select {
		case answer := <-ch:
			point = checkAnswer(answer, quiz.answer, point)
		case <-time.After(timeLimit):
			fmt.Println(" ")
			continue
		}
	}

	return point
}

func getQuizzes(r io.Reader) Quizzes {
	csvReader := csv.NewReader(r)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Failed to parse csv", err)
	}

	quizzes := Quizzes{}
	for _, record := range records {
		quiz := Quiz{
			question: record[0],
			answer:   record[1],
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes
}

func shuffleQuizzes(quizzes Quizzes) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizzes), func(i, j int) { quizzes[i], quizzes[j] = quizzes[j], quizzes[i] })
}
