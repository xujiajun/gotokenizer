package gotokenizer

// WordFilter defines defines interface of WordFilter
type WordFilter interface {
	Filter(text string) bool
}
