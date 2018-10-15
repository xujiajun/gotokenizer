package gotokenizer

var (
	dictZhPath                = "data/zh/dict.txt"
	bigramDictPath            = "data/zh/bigram.txt"
	stopTokensDictPath        = "data/zh/stop_tokens.txt"
	zhText                    = "中华人民共和国万岁万岁万万岁"
	textForBiDirectionalMatch = "这几块地面积还真不小"
	sep                       = "/"
	errorFormat               = "returned unexpected result: got %v want %v"
)
