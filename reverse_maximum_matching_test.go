package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

func TestReverseMaxMatch_Get_With_mixtureText(t *testing.T) {
	rmm := NewReverseMaxMatch(dictZhPath)
	expectedResult := "g/otokenizer/是/一款/基于/字典/和/Bigram/模型/纯/go/语言/编写/的/分词器/，/支持/6/种/分词/算法/。/支持/stopToken/过滤/和/自定义/word/过滤/功能/。"
	rmm.LoadDict()
	result, err := rmm.Get(mixtureText)
	checkErr(err, t)
	reality := strings.Join(result, "/")
	if reality != expectedResult {
		t.Errorf(errorFormat,
			reality, expectedResult)
	}
}

func TestReverseMaxMatch_Get(t *testing.T) {
	expected := "中华人民共和国/万岁/万岁/万万岁"
	rmm := NewReverseMaxMatch(dictZhPath)
	rmm.LoadDict()
	result, err := rmm.Get(zhText)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	if reality != expected {
		t.Errorf(errorFormat,
			reality, expected)
	}
}

func TestReverseMaxMatch_GetFrequency(t *testing.T) {
	expected := "中华人民共和国/万岁/万岁/万万岁"

	rmm := NewReverseMaxMatch(dictZhPath)
	rmm.LoadDict()
	result := strings.Split(expected, sep)
	expectedMap := GetFrequency(result)

	reality, err := rmm.GetFrequency(zhText)
	checkErr(err, t)

	if !reflect.DeepEqual(expectedMap, reality) {
		t.Errorf(errorFormat,
			reality, expected)
	}
}

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
