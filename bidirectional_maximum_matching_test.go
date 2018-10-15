package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForBiDirectionalMaxMatch = "这/几块/地/面积/还/真/不小"
	bdmm                             = NewBiDirectionalMaxMatch(dictZhPath, bigramDictPath)
)

func TestBiDirectionalMaxMatch_Get(t *testing.T) {
	bdmm.LoadDict()

	result, err := bdmm.Get(textForBiDirectionalMatch)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	if reality != expectedForBiDirectionalMaxMatch {
		t.Errorf(errorFormat,
			reality, expectedForBiDirectionalMaxMatch)

	}

	if bdmm.RMMScore < bdmm.MMScore {
		t.Errorf("returned unexpected result: FmmScore %v is bigger want bdmm.BmmScore %v is bigger", bdmm.MMScore, bdmm.RMMScore)
	}

	if bdmm.RMMScore == bdmm.MMScore {
		t.Errorf("returned unexpected result: FmmScore %v is equal to bdmm.BmmScore. bdmm.BmmScore %v is bigger", bdmm.MMScore, bdmm.RMMScore)
	}

}

func TestBiDirectionalMaxMatch_GetFrequency(t *testing.T) {
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
