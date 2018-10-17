package gotokenizer

import (
	"unicode"
)

// NumAndLetterWordFilter that implements the WordFilter interface
type NumAndLetterWordFilter struct {
}

// Filter that implements the WordFilter interface
func (nlFilter *NumAndLetterWordFilter) Filter(text string) bool {
	for _, r := range []rune(text) {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return false
		} else if unicode.IsSpace(r) {
			return false
		} else if unicode.IsControl(r) {
			return false
		} else if unicode.IsPunct(r) {
			return false
		}
	}

	return true
}
