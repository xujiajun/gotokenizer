package go_tokenizer

import (
	"os"
	"io"
	"bufio"
)

type StopTokens struct {
	path string
	records map[string]bool
	IsLoaded bool
}

func NewStopTokens() *StopTokens {
	return &StopTokens{
		records: make(map[string]bool),
	}
}

func (st *StopTokens)Load(path string) error  {
	if st.IsLoaded {
		return nil
	}

	fi, err := os.Open(path)
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
		st.records[string(a)] = true
	}
	st.records[" "]= true
	st.IsLoaded = true

	return nil
}

func (st *StopTokens) IsStopToken(token string) bool {
	_, found := st.records[token]
	return found
}
