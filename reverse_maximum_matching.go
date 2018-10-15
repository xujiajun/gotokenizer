package go_tokenizer

import (
	"strings"
)

type ReverseMaxMatch struct {
	dict     *Dict
	dictPath string
}

func NewReverseMaxMatch(dictPath string) *ReverseMaxMatch {
	return &ReverseMaxMatch{
		dictPath: dictPath,
	}
}

func (rmm *ReverseMaxMatch) LoadDict() error {
	rmm.dict = NewDict(rmm.dictPath)
	return rmm.dict.Load()
}

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

func (rmm *ReverseMaxMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := rmm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
