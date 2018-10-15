package gotokenizer

import (
	"strings"
)

type MaxMatch struct {
	dict     *Dict
	dictPath string
}

func NewMaxMatch(dictPath string) *MaxMatch {
	return &MaxMatch{
		dictPath: dictPath,
	}
}

func (mm *MaxMatch) LoadDict() error {
	mm.dict = NewDict(mm.dictPath)
	return mm.dict.Load()
}

func (mm *MaxMatch) Get(text string) ([]string, error) {

	CheckDictIsLoaded(mm.dict)

	var result []string

	startLen := mm.dict.maxLen
	text = strings.Trim(text, " ")

	for len([]rune(text)) > 0 {

		if len([]rune(text)) < startLen {
			startLen = len([]rune(text))
		}

		word := string([]rune(text)[0:startLen])

		isFind := false
		for !isFind {

			if len([]rune(word)) == 1 {
				break
			}

			if _, ok := mm.dict.Records[word]; !ok {
				word = string([]rune(word)[0 : len([]rune(word))-1])
			} else {
				isFind = true
			}
		}

		result = append(result, word)
		text = string([]rune(text)[len([]rune(word)):])
	}

	return result, nil
}

func (mm *MaxMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := mm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
