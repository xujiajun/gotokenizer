# gotokenizer [![GoDoc](https://godoc.org/github.com/xujiajun/gotokenizer?status.svg)](https://godoc.org/github.com/xujiajun/gotokenizer) <a href="https://travis-ci.org/xujiajun/gotokenizer"><img src="https://travis-ci.org/xujiajun/gotokenizer.svg?branch=master" alt="Build Status"></a> [![Coverage Status](https://coveralls.io/repos/github/xujiajun/gotokenizer/badge.svg?branch=master)](https://coveralls.io/github/xujiajun/gotokenizer?branch=master) [![License](https://img.shields.io/badge/license-Apache2.0-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)
A tokenizer for Go. (Now only support chinese segmentation)

## Motivation

I wanted a simple tokenizer that has no unnecessary overhead using the standard library only, following good practices and well tested code.

## Features

* Support Maximum Matching Method
* Support Minimum Matching Method
* Support Reverse Maximum Matching
* Support Reverse Minimum Matching
* Support Bidirectional Maximum Matching
* Support Bidirectional Minimum Matching
* Support using Stop Tokens

## Installation

```
go get -u github.com/xujiajun/gotokenizer
```

## Usage

```
package main

import (
	"fmt"
	"github.com/xujiajun/gotokenizer"
)

func main() {
	text := "中华人民共和国万岁万岁万万岁"
	dictPath := "/Users/xujiajun/go/src/github.com/xujiajun/gotokenizer/data/zh/dict.txt" // use your dict
	mm := gotokenizer.NewMaxMatch(dictPath)
	mm.LoadDict()
	result, err := mm.Get(text)
	fmt.Println(result, err) //result: [中华人民共和国 万岁 万岁 万万岁] <nil>
}

```

> More examples see tests

## Contributing

If you'd like to help out with the project. You can put up a Pull Request.


## Author

* [xujiajun](https://github.com/xujiajun)

## License

The gotokenizer is open-sourced software licensed under the [Apache-2.0](https://opensource.org/licenses/Apache-2.0)

## Acknowledgements

This package is inspired by the following:

https://github.com/ysc/word
