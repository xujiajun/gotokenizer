package gotokenizer

import (
	"strings"
	"testing"
)

var testText = []string{"go", "for", "tokenizer", "a", "is", "go-tokenizer"}

func TestReverse(t *testing.T) {
	expectedText := "go-tokenizer is a tokenizer for go"
	reality := strings.Join(Reverse(testText), " ")
	if reality != expectedText {
		t.Errorf(errorFormat, reality, expected)
	}
}

func TestCheckDictIsLoaded(t *testing.T) {
	if CheckDictIsLoaded(nil) == nil {
		t.Errorf(errorFormat, CheckDictIsLoaded(nil), nil)
	}
}
