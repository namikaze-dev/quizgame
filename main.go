package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func PlayGame(o io.Writer, i io.Reader, questions []ParsedQuestion) {
	var corrects int
	var input = bufio.NewScanner(i)
	for _, question := range questions {
		answer := playNextLine(o, input, question)
		if answer == question.Answer {
			corrects += 1
		}
	}
	fmt.Fprintf(o, "\nCORRECT: %v\nTOTAL: %v", corrects, len(questions))
}

func playNextLine(o io.Writer, scn *bufio.Scanner, question ParsedQuestion) string {
	fmt.Fprintf(o, "QUESTION: %v\n", question.Question)
	answer := readNextAnswer(scn)
	fmt.Fprintf(o, "ANSWER: %v\n", answer)
	return answer
}

func readNextAnswer(scn *bufio.Scanner) string {
	scn.Scan()
	return strings.TrimSpace(strings.ToLower(scn.Text()))
}
