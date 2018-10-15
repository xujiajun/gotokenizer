package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForNewBackwardMinMatch = "中华/人民/共和国/万岁/万岁/万/万岁"
	bminm                          = NewReverseMinMatch(dictZhPath)
)

func TestReverseMinMatch_Get(t *testing.T) {
	bminm.LoadDict()
	result, err := bminm.Get(zhText)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	if reality != expectedForNewBackwardMinMatch {
		t.Errorf(errorFormat,
			reality, expectedForNewBackwardMinMatch)

	}
}

func TestReverseMinMatch_GetFrequency(t *testing.T) {
	bminm.LoadDict()
	result := strings.Split(expectedForNewBackwardMinMatch, sep)
	expected := GetFrequency(result)

	reality, err := bminm.GetFrequency(zhText)
	checkErr(err, t)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expected)
	}
}
