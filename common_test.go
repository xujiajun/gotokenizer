package gotokenizer

var (
	dictZhPath                = "data/zh/dict.txt"
	bigramDictPath            = "data/zh/bigram.txt"
	stopTokensDictPath        = "data/zh/stop_tokens.txt"
	zhText                    = "中华人民共和国万岁万岁万万岁"
	mixtureText               = "gotokenizer是一款基于字典和Bigram模型纯go语言编写的分词器，支持6种分词算法。支持stopToken过滤和自定义word过滤功能。"
	textForBiDirectionalMatch = "这几块地面积还真不小"
	sep                       = "/"
	errorFormat               = "returned unexpected result: got %v want %v"
)
