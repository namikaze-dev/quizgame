package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

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
	return answer
}

func readNextAnswer(scn *bufio.Scanner) string {
	scn.Scan()
	return strings.TrimSpace(strings.ToLower(scn.Text()))
}
