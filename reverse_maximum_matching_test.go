package go_tokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	rmm      = NewReverseMaxMatch(dictZhPath)
	expected = "中华人民共和国/万岁/万岁/万万岁"
)

func TestReverseMaxMatch_Get(t *testing.T) {
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
	rmm.LoadDict()
	result := strings.Split(expected, sep)
	expected := GetFrequency(result)

	reality, err := rmm.GetFrequency(zhText)
	checkErr(err, t)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expected)
	}
}

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
