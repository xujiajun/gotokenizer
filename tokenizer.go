package gotokenizer

type Tokenizer interface {
	GetFrequency(text string) (map[string]int, error)
	Get(text string) ([]string, error)
	LoadDict() error
}
