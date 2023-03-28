package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Options struct {
	Source  string
	Timeout int
	Shuffle bool
}

func main() {
	var options Options
	flag.StringVar(&options.Source, "file", "problems.csv", "source csv file for questions")
	flag.IntVar(&options.Timeout, "timeout", 30, "quiz time limit")
	flag.BoolVar(&options.Shuffle, "shuffle", false, "shuffle questions")
	flag.Parse()

	source := openFile(options.Source)
	questions := ParseQuestionsFromReader(source)
	if options.Shuffle {
		shuffle(questions)
	}

	PlayGame(os.Stdout, os.Stdin, questions, options)
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("quizgame:", err)
		os.Exit(-1)
	}
	return f
}

func shuffle(questions []ParsedQuestion) {
	rand.Seed(time.Now().UnixMilli())
	for i := range questions {
		j := rand.Intn(len(questions))
		questions[i], questions[j] = questions[j], questions[i]
	}
}
