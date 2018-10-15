package go_tokenizer

var (
	dictZhPath                = "data/zh/dict.txt"
	dictEnPath                = "data/en/dict.txt"
	bigramDictPath            = "data/zh/bigram.txt"
	stopTokensDictPath        = "data/stop_tokens.txt"
	zhText                    = "中华人民共和国万岁万岁万万岁"
	textForBiDirectionalMatch = "这几块地面积还真不小"
	sep                       = "/"
	errorFormat               = "returned unexpected result: got %v want %v"
)
