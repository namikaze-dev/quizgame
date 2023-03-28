package main

import (
	"bufio"
	"io"
	"strings"
)

type ParsedQuestion struct {
	Question, Answer string
}

func ParseQuestionsFromReader(r io.Reader) []ParsedQuestion {
	var questions []ParsedQuestion
	scn := bufio.NewScanner(r)
	for scn.Scan() {
		question, ok := parseNextLine(scn)
		if ok {
			questions = append(questions, question)
		}
	}
	return questions
}

func parseNextLine(scn *bufio.Scanner) (ParsedQuestion, bool) {
	line := scn.Text()
	sepIdx := strings.LastIndexByte(line, ',')

	if sepIdx == -1 {
		return ParsedQuestion{}, false
	}
	
	return ParsedQuestion{
		Question: strings.TrimSpace(strings.ToLower(line[:sepIdx])),
		Answer:   strings.TrimSpace(strings.ToLower(line[sepIdx+1:])),
	}, true
}
