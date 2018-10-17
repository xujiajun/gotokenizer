package gotokenizer

type WordFilter interface {
	Filter(text string) bool
}
