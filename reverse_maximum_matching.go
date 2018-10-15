package gotokenizer

import (
	"strings"
)

// ReverseMaxMatch records dict and dictPath
type ReverseMaxMatch struct {
	dict     *Dict
	dictPath string
}

// NewReverseMaxMatch returns a newly initialized ReverseMaxMatch object
func NewReverseMaxMatch(dictPath string) *ReverseMaxMatch {
	return &ReverseMaxMatch{
		dictPath: dictPath,
	}
}

// LoadDict loads dict that implements the Tokenizer interface
func (rmm *ReverseMaxMatch) LoadDict() error {
	rmm.dict = NewDict(rmm.dictPath)
	return rmm.dict.Load()
}

// Get returns segmentation that implements the Tokenizer interface
func (rmm *ReverseMaxMatch) Get(text string) ([]string, error) {

	CheckDictIsLoaded(rmm.dict)

	var result []string

	startLen := rmm.dict.maxLen

	text = strings.Trim(text, " ")

	for len([]rune(text)) > 0 {

		if len([]rune(text)) < startLen {
			startLen = len([]rune(text))
		}

		word := string([]rune(text)[len([]rune(text))-startLen:])

		isFind := false

		for !isFind {
			if len([]rune(word)) == 1 {
				break
			}

			if _, ok := rmm.dict.Records[word]; !ok {
				word = string([]rune(word)[1:])
			} else {
				isFind = true
			}
		}

		result = append(result, word)
		text = string([]rune(text)[0 : len([]rune(text))-len([]rune(word))])
	}

	return Reverse(result), nil
}

// GetFrequency returns token frequency that implements the Tokenizer interface
func (rmm *ReverseMaxMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := rmm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
