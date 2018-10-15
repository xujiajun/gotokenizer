package go_tokenizer

import (
	"strings"
)

type ReverseMinMatch struct {
	dict     *Dict
	dictPath string
}

func NewReverseMinMatch(dictPath string) *ReverseMinMatch {
	return &ReverseMinMatch{
		dictPath: dictPath,
	}
}

func (rmm *ReverseMinMatch) LoadDict() error {
	rmm.dict = NewDict(rmm.dictPath)
	return rmm.dict.Load()
}

func (rmm *ReverseMinMatch) Get(text string) ([]string, error) {

	CheckDictIsLoaded(rmm.dict)

	var result []string

	startLen := 2
	text = strings.Trim(text, " ")

	for startLen <= len([]rune(text)) {

		word := string([]rune(text)[len([]rune(text))-startLen:])

		isFind := false

		for !isFind {
			if len([]rune(text)) == 2 {
				word = string([]rune(text))
			} else {
				if startLen == len([]rune(text))-1 {
					startLen = 1
					word = string([]rune(text)[len([]rune(text))-startLen:])
					break
				}
			}

			if _, ok := rmm.dict.Records[word]; !ok {
				startLen++
				if startLen > len([]rune(text)) {
					break
				}
				word = string([]rune(text)[len([]rune(text))-startLen:])
			} else {
				startLen = 2
				isFind = true
			}

		}

		result = append(result, word)
		text = string([]rune(text)[0 : len([]rune(text))-len([]rune(word))])
	}

	return Reverse(result), nil
}

func (rmm *ReverseMinMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := rmm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
