package gotokenizer

import (
	"strings"
)

// MinMatch records dict and dictPath
type MinMatch struct {
	dict     *Dict
	dictPath string
}

// NewMinMatch returns a newly initialized MinMatch object
func NewMinMatch(dictPath string) *MinMatch {
	return &MinMatch{
		dictPath: dictPath,
	}
}

// LoadDict loads dict that implements the Tokenizer interface
func (mm *MinMatch) LoadDict() error {
	mm.dict = NewDict(mm.dictPath)
	return mm.dict.Load()
}

// Get returns segmentation that implements the Tokenizer interface
func (mm *MinMatch) Get(text string) ([]string, error) {
	CheckDictIsLoaded(mm.dict)

	var result []string

	startLen := 2

	text = strings.Trim(text, " ")

	for len([]rune(text)) > 0 {

		word := string([]rune(text)[0:startLen])

		isFind := false
		for !isFind {

			if startLen == mm.dict.maxLen || startLen == len([]rune(text)) || len([]rune(text)) == 1 {
				startLen = 1
				word = string([]rune(text)[0:startLen])
				break
			}

			if _, ok := mm.dict.Records[word]; !ok {

				startLen++

				if startLen > len([]rune(text)) {
					break
				}

				word = string([]rune(text)[0:startLen])
			} else {
				startLen = 2
				isFind = true
			}

		}

		result = append(result, word)

		text = string([]rune(text)[len([]rune(word)):])
	}

	return result, nil
}

// GetFrequency returns token frequency that implements the Tokenizer interface
func (mm *MinMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := mm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
