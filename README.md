# gotokenizer [![GoDoc](https://godoc.org/github.com/xujiajun/gotokenizer?status.svg)](https://godoc.org/github.com/xujiajun/gotokenizer) <a href="https://travis-ci.org/xujiajun/gotokenizer"><img src="https://travis-ci.org/xujiajun/gotokenizer.svg?branch=master" alt="Build Status"></a> [![Coverage Status](https://coveralls.io/repos/github/xujiajun/gotokenizer/badge.svg?branch=master)](https://coveralls.io/github/xujiajun/gotokenizer?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/xujiajun/gotokenizer)](https://goreportcard.com/report/github.com/xujiajun/gotokenizer) [![License](https://img.shields.io/badge/license-Apache2.0-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go#natural-language-processing) 
A tokenizer based on the dictionary and Bigram language models for Golang.  (Now only support chinese segmentation)

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
* Support Custom word Filter

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
	text := "gotokenizer是一款基于字典和Bigram模型纯go语言编写的分词器，支持6种分词算法。支持stopToken过滤和自定义word过滤功能。"

	dictPath := "/Users/xujiajun/go/src/github.com/xujiajun/gotokenizer/data/zh/dict.txt"
	// NewMaxMatch default wordFilter is NumAndLetterWordFilter
	mm := gotokenizer.NewMaxMatch(dictPath)
	// load dict
	mm.LoadDict()

	fmt.Println(mm.Get(text)) //[gotokenizer 是 一款 基于 字典 和 Bigram 模型 纯 go 语言 编写 的 分词器 ， 支持 6 种 分词 算法 。 支持 stopToken 过滤 和 自定义 word 过滤 功能 。] <nil>

	// enabled filter stop tokens 
	mm.EnabledFilterStopToken = true
	mm.StopTokens = gotokenizer.NewStopTokens()
	stopTokenDicPath := "/Users/xujiajun/go/src/github.com/xujiajun/gotokenizer/data/zh/stop_tokens.txt"
	mm.StopTokens.Load(stopTokenDicPath)

	fmt.Println(mm.Get(text)) //[gotokenizer 一款 字典 Bigram 模型 go 语言 编写 分词器 支持 6 种 分词 算法 支持 stopToken 过滤 自定义 word 过滤 功能] <nil>
	fmt.Println(mm.GetFrequency(text)) //map[6:1 种:1 算法:1 过滤:2 支持:2 Bigram:1 模型:1 编写:1 gotokenizer:1 go:1 分词器:1 分词:1 word:1 功能:1 一款:1 语言:1 stopToken:1 自定义:1 字典:1] <nil>

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
