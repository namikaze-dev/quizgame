package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	source string
}

func main() {
	var options Options
	flag.StringVar(&options.source, "file", "problems.csv", "source csv file for questions")
	flag.Parse()

	source := openFile(options.source)
	questions := ParseQuestionsFromReader(source)
	PlayGame(os.Stdout, os.Stdin, questions)
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("quizgame:", err)
		os.Exit(-1)
	}
	return f
}
