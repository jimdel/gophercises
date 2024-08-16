package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "quiz.csv", "CSV file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal("Failed to open the CSV file: ", *csvFilename)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to read the CSV file")
	}

	problems := ParseQuiz(lines)
	score := RunQuiz(os.Stdin, problems, timeLimit)
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func ParseQuiz(lines [][]string) []problem {
	formattedQuiz := make([]problem, len(lines))
	for idx, line := range lines {
		formattedQuiz[idx] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return formattedQuiz
}

// We use an ioReader as the input to make testing easier
func RunQuiz(input io.Reader, quiz []problem, timeLimit *int) int {
	score := 0
	scanner := bufio.NewScanner(input)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for idx, problem := range quiz {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.q)
		answerChan := make(chan string)
		go func() {
			if scanner.Scan() {
				answer := scanner.Text()
				answerChan <- answer
			}
		}()
		select {
		case <-timer.C:
			return score
		case answer := <-answerChan:
			if answer == problem.a {
				score++
			}
		}
	}
	return score
}
