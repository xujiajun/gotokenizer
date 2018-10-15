package gotokenizer

import (
	"bufio"
	"io"
	"os"
)

// StopTokens records paths and records
type StopTokens struct {
	path     string
	records  map[string]bool
	IsLoaded bool
}

// NewStopTokens returns a newly initialized StopTokens object
func NewStopTokens() *StopTokens {
	return &StopTokens{
		records: make(map[string]bool),
	}
}

// Load that loads StopToken dict
func (st *StopTokens) Load(path string) error {
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
	st.records[" "] = true
	st.IsLoaded = true

	return nil
}

// IsStopToken returns if token is a token
func (st *StopTokens) IsStopToken(token string) bool {
	_, found := st.records[token]
	return found
}
