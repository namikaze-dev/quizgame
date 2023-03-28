package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"time"
)

func PlayGame(o io.Writer, i io.Reader, questions []ParsedQuestion, options Options) {
	var scn = bufio.NewScanner(i)
	// scan once as starting trigger
	scn.Scan()

	var ctx, cancel = context.WithTimeout(context.Background(), time.Duration(options.Timeout)*time.Second)
	defer cancel()

	var corrects int
	for _, question := range questions {
		fmt.Fprintf(o, "QUESTION: %v\n", question.Question)
		answer, err := readNextAnswer(ctx, scn)

		if err != nil {
			fmt.Fprintln(o, "quizgame: time limit exceeded")
			break
		}

		if answer == question.Answer {
			corrects += 1
		}
	}
	fmt.Fprintf(o, "\nCORRECT: %v\nTOTAL: %v", corrects, len(questions))
}

func readNextAnswer(ctx context.Context, scn *bufio.Scanner) (string, error) {
	var result = make(chan string)

	go func() {
		scn.Scan()
		result <- strings.TrimSpace(strings.ToLower(scn.Text()))
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case ans := <-result:
		return ans, nil
	}
}
