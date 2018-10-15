package gotokenizer

import (
	"bufio"
	"github.com/xujiajun/utils/strconv2"
	"io"
	"os"
	"strings"
)

// BigramDict records dictPath and bigram records
type BigramDict struct {
	dictPath string
	isLoaded bool
	maxF     int
	records  map[string]int
}

// NewBigramDict returns a newly initialized BigramDict object
func NewBigramDict(dictPath string) *BigramDict {
	return &BigramDict{
		dictPath: dictPath,
		records:  make(map[string]int),
	}
}

// Load returns Bigram Dict records
func (bd *BigramDict) Load() error {
	if bd.isLoaded {
		return nil
	}

	fi, err := os.Open(bd.dictPath)
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

		f := res[1]
		fi, _ := strconv2.StrToInt(f)

		strKey := res[0]
		bd.records[strKey] = fi

		if bd.maxF < fi {
			bd.maxF = fi
		}
	}

	bd.isLoaded = true

	return nil
}
