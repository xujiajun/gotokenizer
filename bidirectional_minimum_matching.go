package gotokenizer

import (
	"strings"
)

// BiDirectionalMinMatch records dict and bigramDic etc.
type BiDirectionalMinMatch struct {
	dict           *Dict
	dictPath       string
	bigramDictPath string
	bigramDic      *BigramDict
	MMScore        float64
	RMMScore       float64
	MM             *MinMatch
	RMM            *ReverseMinMatch
}

// NewBiDirectionalMinMatch returns a newly initialized BiDirectionalMinMatch object
func NewBiDirectionalMinMatch(dictPath, bigramDictPath string) *BiDirectionalMinMatch {
	return &BiDirectionalMinMatch{
		dictPath:       dictPath,
		dict:           NewDict(dictPath),
		bigramDictPath: bigramDictPath,
		bigramDic:      NewBigramDict(bigramDictPath),
		MM:             NewMinMatch(dictPath),
		RMM:            NewReverseMinMatch(dictPath),
	}
}

// LoadDict load dict and bigramDic that implements the Tokenizer interface
func (bdmm *BiDirectionalMinMatch) LoadDict() error {
	bdmm.dict.Load()
	bdmm.bigramDic.Load()
	return nil
}

// Get returns segmentation that implements the Tokenizer interface
func (bdmm *BiDirectionalMinMatch) Get(text string) ([]string, error) {
	text = strings.Trim(text, " ")

	bdmm.MM.dict = bdmm.dict
	mmResult, _ := bdmm.MM.Get(text)

	bdmm.RMM.dict = bdmm.dict
	rmmResult, _ := bdmm.RMM.Get(text)

	for i := 0; i < len(mmResult)-1; i++ {
		key := mmResult[i] + ":" + mmResult[i+1]

		if val, ok := bdmm.bigramDic.records[key]; ok {
			score := float64(val) / float64(bdmm.bigramDic.maxF)
			bdmm.MMScore += score
		}
	}

	for i := 0; i < len(rmmResult)-1; i++ {
		key := rmmResult[i] + ":" + rmmResult[i+1]

		if val, ok := bdmm.bigramDic.records[key]; ok {
			score := float64(val) / float64(bdmm.bigramDic.maxF)
			bdmm.RMMScore += score
		}
	}

	if bdmm.MMScore > bdmm.RMMScore {
		return mmResult, nil
	}

	if bdmm.MMScore < bdmm.RMMScore {
		return rmmResult, nil
	}

	if bdmm.MMScore == bdmm.RMMScore {
		return mmResult, nil
	}

	return nil, nil
}

// GetFrequency returns token frequency that implements the Tokenizer interface
func (bdmm *BiDirectionalMinMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := bdmm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
