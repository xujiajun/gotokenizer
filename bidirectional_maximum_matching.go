package gotokenizer

import (
	"strings"
)

type BiDirectionalMaxMatch struct {
	dict           *Dict
	dictPath       string
	bigramDictPath string
	bigramDic      *BigramDict
	MMScore        float64
	RMMScore       float64
	MM             *MaxMatch
	RMM            *ReverseMaxMatch
}

func NewBiDirectionalMaxMatch(dictPath, bigramDictPath string) *BiDirectionalMaxMatch {
	return &BiDirectionalMaxMatch{
		dictPath:       dictPath,
		dict:           NewDict(dictPath),
		bigramDictPath: bigramDictPath,
		bigramDic:      NewBigramDict(bigramDictPath),
		MM:             NewMaxMatch(dictPath),
		RMM:            NewReverseMaxMatch(dictPath),
	}
}

func (bdmm *BiDirectionalMaxMatch) LoadDict() error {
	bdmm.dict.Load()
	bdmm.bigramDic.Load()
	return nil
}

func (bdmm *BiDirectionalMaxMatch) Get(text string) ([]string, error) {
	text = strings.Trim(text, " ")

	bdmm.MM.dict = bdmm.dict
	fmmResult, _ := bdmm.MM.Get(text)

	bdmm.RMM.dict = bdmm.dict
	bmmResult, _ := bdmm.RMM.Get(text)

	for i := 0; i < len(fmmResult)-1; i++ {
		key := fmmResult[i] + ":" + fmmResult[i+1]

		if val, ok := bdmm.bigramDic.records[key]; ok {
			score := float64(val) / float64(bdmm.bigramDic.maxF)
			bdmm.MMScore += score
		}
	}

	for i := 0; i < len(bmmResult)-1; i++ {
		key := bmmResult[i] + ":" + bmmResult[i+1]

		if val, ok := bdmm.bigramDic.records[key]; ok {
			score := float64(val) / float64(bdmm.bigramDic.maxF)
			bdmm.RMMScore += score
		}
	}

	if bdmm.MMScore > bdmm.RMMScore {
		return fmmResult, nil
	}

	if bdmm.MMScore < bdmm.RMMScore {
		return bmmResult, nil
	}

	if bdmm.MMScore == bdmm.RMMScore {
		return fmmResult, nil
	}

	return nil, nil
}

func (bdmm *BiDirectionalMaxMatch) GetFrequency(text string) (map[string]int, error) {
	result, err := bdmm.Get(text)

	if err != nil {
		return nil, err
	}

	return GetFrequency(result), nil
}
