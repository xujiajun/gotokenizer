package go_tokenizer

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var DefaultMinTokenLen = 2

type DictRecord struct {
	TF    string
	Token string
	POS   string //part of speech
}

type Dict struct {
	Records     map[string]DictRecord
	minTokenLen int
	DictPath    string
	maxLen      int
	isLoaded    bool
}

func NewDict(dictPath string) *Dict {
	return &Dict{
		Records:     make(map[string]DictRecord),
		DictPath:    dictPath,
		minTokenLen: DefaultMinTokenLen,
		isLoaded:    false,
	}
}

func (dict *Dict) Load() error {
	if dict.isLoaded {
		return errors.New("dict isLoaded")
	}

	fi, err := os.Open(dict.DictPath)
	if err != nil {
		return err
	}

	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		res := strings.Split(string(a), " ")

		var TF,pos string

		token := res[0]

		if len(res) > 1 {
			TF = res[1]
			pos = res[2]
		}

		currLen := len([]rune(token))
		if currLen > dict.maxLen {
			dict.maxLen = currLen
		}


		if len([]rune(token)) >= dict.minTokenLen {
			dict.Records[token] = DictRecord{
				TF:  TF,
				POS: pos,
			}
		}
	}

	dict.isLoaded = true

	return nil
}
