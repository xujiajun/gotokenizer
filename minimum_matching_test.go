package go_tokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForForwardMinMatch = "中华/人民/共和/国/万岁/万岁/万万/岁"
	fmm                        = NewMinMatch(dictZhPath)
)

func TestMinMatch_Get(t *testing.T) {
	fmm.LoadDict()
	result, err := fmm.Get(zhText)
	checkErr(err, t)

	reality := strings.Join(result, sep)

	if reality != expectedForForwardMinMatch {
		t.Errorf(errorFormat,
			reality, expectedForForwardMinMatch)

	}
}

func TestMinMatch_GetFrequency(t *testing.T) {
	fmm.LoadDict()
	result := strings.Split(expectedForForwardMinMatch, sep)
	expected := GetFrequency(result)

	reality, err := fmm.GetFrequency(zhText)
	checkErr(err, t)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expectedForForwardMinMatch)
	}
}
