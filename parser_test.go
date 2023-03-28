package main_test

import (
	"reflect"
	"strings"
	"testing"

	main "github.com/namikaze-dev/quizgame"
)

func TestParseQuestionsFromReader(t *testing.T) {
	rd := strings.NewReader(sample_csv)
	got := main.ParseQuestionsFromReader(rd)
	want := []main.ParsedQuestion{
		{Question: "5+5", Answer: "10"},
		{Question: "8+3", Answer: "11"},
		{Question: "1+2", Answer: "3"},
		{Question: `"whats the fastest animal on land"`, Answer: "cheetah"},
		{Question: `"what 2*2, sir?"`, Answer: "4"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

const sample_csv = `5+5,10
8+3,11
1+2,3

"whats the fastest animal on land"  ,cheetah
"what 2*2, sir?",4`
