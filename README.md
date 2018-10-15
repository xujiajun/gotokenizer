# go-tokenizer
A tokenizer for Go. (Now only support chinese segmentation)

## Usage

```
package main

import (
	"fmt"
	"github.com/xujiajun/go-tokenizer"
)

func main() {
	text := "中华人民共和国万岁万岁万万岁"
	dictPath := "/Users/xujiajun/go/src/github.com/xujiajun/go-tokenizer/data/zh/dict.txt" // use your dict
	mm := go_tokenizer.NewMaxMatch(dictPath)
	mm.LoadDict()
	result, err := mm.Get(text)
	fmt.Println(result, err) //result: [中华人民共和国 万岁 万岁 万万岁] <nil>
}

```

> More see tests

## Contributing

If you'd like to help out with the project. You can put up a Pull Request.

## Author

* [xujiajun](https://github.com/xujiajun)

## License

The go-tokenizer is open-sourced software licensed under the [Apache-2.0](https://opensource.org/licenses/Apache-2.0)
