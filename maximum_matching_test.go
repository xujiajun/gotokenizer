package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForForwardMaxMatch = "中华人民共和国/万岁/万岁/万万岁"
	mm                         = NewMaxMatch(dictZhPath)
)

func TestMaxMatch_Get_With_mixtureText(t *testing.T) {
	expectedResult := "gotokenizer/是/一款/基于/字典/和/Bigram/模型/纯/go/语言/编写/的/分词器/，/支持/6/种/分词/算法/。/支持/stopToken/过滤/和/自定义/word/过滤/功能/。"
	mm.LoadDict()
	result, err := mm.Get(mixtureText)
	checkErr(err, t)
	reality := strings.Join(result, "/")
	if reality != expectedResult {
		t.Errorf(errorFormat,
			reality, expectedResult)
	}
}

func TestMaxMatch_Get_enabledStopToken(t *testing.T) {
	mm.LoadDict()
	mm.EnabledFilterStopToken = true
	mm.StopTokens = NewStopTokens()
	mm.StopTokens.Load(stopTokensDictPath)
	result, err := mm.Get(mixtureText)
	checkErr(err, t)

	expectedResult := "gotokenizer/一款/字典/Bigram/模型/go/语言/编写/分词器/支持/6/种/分词/算法/支持/stopToken/过滤/自定义/word/过滤/功能"
	reality := strings.Join(result, "/")
	if reality != expectedResult {
		t.Errorf(errorFormat,
			reality, expectedResult)
	}
}

func TestMaxMatch_Get(t *testing.T) {
	mm.LoadDict()
	result, err := mm.Get(zhText)
	checkErr(err, t)

	reality := strings.Join(result, "/")
	if reality != expectedForForwardMaxMatch {
		t.Errorf(errorFormat,
			reality, expectedForForwardMaxMatch)

	}
}

func TestMaxMatch_GetFrequency(t *testing.T) {
	mm.LoadDict()

	result := strings.Split(expectedForForwardMaxMatch, "/")
	expected := GetFrequency(result)

	reality, err := mm.GetFrequency(zhText)
	checkErr(err, t)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expectedForForwardMaxMatch)
	}
}
