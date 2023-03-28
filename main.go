package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Source string
	Timeout int
}

func main() {
	var options Options
	flag.StringVar(&options.Source, "file", "problems.csv", "source csv file for questions")
	flag.IntVar(&options.Timeout, "timeout", 30, "quiz time limit")
	flag.Parse()

	source := openFile(options.Source)
	questions := ParseQuestionsFromReader(source)
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
