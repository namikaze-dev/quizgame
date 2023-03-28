package main_test

import (
	"bytes"
	"strings"
	"testing"

	main "github.com/namikaze-dev/quizgame"
)

func TestPlayGame(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		parsedQuestions := main.ParseQuestionsFromReader(strings.NewReader(sample_csv))
		output := &bytes.Buffer{}
		input := strings.NewReader(sample_csv_answers)
		main.PlayGame(output, input, parsedQuestions, main.Options{Timeout: 30})

		got := output.String()
		want := `QUESTION: 5+5
QUESTION: 8+3
QUESTION: 1+2
QUESTION: "whats the fastest animal on land"
QUESTION: "what 2*2, sir?"

CORRECT: 3
TOTAL: 5`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
	t.Run("timeout failure", func(t *testing.T) {
		parsedQuestions := main.ParseQuestionsFromReader(strings.NewReader(sample_csv))
		output := &bytes.Buffer{}
		input := strings.NewReader(sample_csv_answers)
		main.PlayGame(output, input, parsedQuestions, main.Options{Timeout: 0})

		got := output.String()
		want := `QUESTION: 5+5
quizgame: time limit exceeded

CORRECT: 0
TOTAL: 5`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

const sample_csv_answers = `10
11
5
cheetah
8`
