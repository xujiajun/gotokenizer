package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForBiDirectionalMaxMatch = "这/几块/地/面积/还/真/不小"
)

func TestBiDirectionalMaxMatch_Get(t *testing.T) {
	bdmm := NewBiDirectionalMaxMatch(dictZhPath, bigramDictPath)
	bdmm.LoadDict()

	result, err := bdmm.Get(textForBiDirectionalMatch)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	if reality != expectedForBiDirectionalMaxMatch {
		t.Errorf(errorFormat,
			reality, expectedForBiDirectionalMaxMatch)

	}

	if bdmm.RMMScore < bdmm.MMScore {
		t.Errorf("returned unexpected result: MMScore %v is bigger want bdmm.RMMScore %v is bigger", bdmm.MMScore, bdmm.RMMScore)
	}

	if bdmm.RMMScore == bdmm.MMScore {
		t.Errorf("returned unexpected result: MMScore %v is equal to bdmm.RMMScore. bdmm.RMMScore %v is bigger", bdmm.MMScore, bdmm.RMMScore)
	}

}

func TestBiDirectionalMaxMatch_GetFrequency(t *testing.T) {
	bdmm := NewBiDirectionalMaxMatch(dictZhPath, bigramDictPath)
	bdmm.LoadDict()

	reality, err := bdmm.GetFrequency(textForBiDirectionalMatch)
	checkErr(err, t)

	result := strings.Split(expectedForBiDirectionalMaxMatch, sep)
	expected := GetFrequency(result)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expected)
	}
}

func TestBiDirectionalMaxMatch_Get_With_MixtureText(t *testing.T) {
	bdmm := NewBiDirectionalMaxMatch(dictZhPath, bigramDictPath)
	bdmm.LoadDict()

	result, err := bdmm.Get(mixtureText)
	checkErr(err, t)

	reality := strings.Join(result, sep)

	expected := "gotokenizer/是/一款/基于/字典/和/Bigram/模型/纯/go/语言/编写/的/分词器/，/支持/6/种/分词/算法/。/支持/stopToken/过滤/和/自定义/word/过滤/功能/。"
	if reality != expected {
		t.Errorf(errorFormat,
			reality, expected)
	}
}

func TestBiDirectionalMaxMatch_Get_enabledFilterStopToken(t *testing.T) {
	bdmm := NewBiDirectionalMaxMatch(dictZhPath, bigramDictPath)
	bdmm.LoadDict()
	stopTokens := NewStopTokens()
	stopTokens.Load(stopTokensDictPath)

	bdmm.MM.EnabledFilterStopToken = true
	bdmm.MM.StopTokens = stopTokens
	bdmm.RMM.EnabledFilterStopToken = true
	bdmm.RMM.StopTokens = stopTokens

	result, err := bdmm.Get(mixtureText)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	expected := "gotokenizer/一款/字典/Bigram/模型/go/语言/编写/分词器/支持/6/种/分词/算法/支持/stopToken/过滤/自定义/word/过滤/功能"
	if reality != expected {
		t.Errorf(errorFormat,
			reality, expected)
	}
}
