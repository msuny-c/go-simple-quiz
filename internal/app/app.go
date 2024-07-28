package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"quiz/internal/csv"
	"quiz/internal/model"
	"strings"
	"time"
)

func Run() {

	csvFilepath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'.")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var quizzes = csv.ReadQuizzes(file)

	channel := make(chan struct{})
	var scored int

	go ask(channel, quizzes, &scored)

	select {
	case <- time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Printf("\nYou scored %d out of %d.", scored, len(quizzes))
	case <- channel:
		fmt.Printf("You scored %d out of %d.", scored, len(quizzes))
	}
}

func ask(channel chan struct{}, quizzes []model.Quiz, scored *int) {
	for i, quiz := range quizzes {
		var input string
		fmt.Printf("Problem #%d: %s = ", i + 1, quiz.Question)
		fmt.Scan(&input)
		if compare(input, quiz.Answer) {
			*scored++
		}
	}
	close(channel)
}

func compare(a, b string) bool {
	return strings.TrimSpace(a) == strings.TrimSpace(b)
}