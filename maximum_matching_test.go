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
