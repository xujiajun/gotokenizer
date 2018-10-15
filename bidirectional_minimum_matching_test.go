package gotokenizer

import (
	"reflect"
	"strings"
	"testing"
)

var (
	expectedForBiDirectionalMinMatch = "这几/块/地/面积/还/真/不小"
	bdminm                           = NewBiDirectionalMinMatch(dictZhPath, bigramDictPath)
)

func TestBiDirectionalMinMatch_Get(t *testing.T) {
	bdminm.LoadDict()

	result, err := bdminm.Get(textForBiDirectionalMatch)
	checkErr(err, t)

	reality := strings.Join(result, sep)
	if reality != expectedForBiDirectionalMinMatch {
		t.Errorf(errorFormat,
			reality, expectedForBiDirectionalMaxMatch)

	}

	if bdminm.RMMScore < bdminm.MMScore {
		t.Errorf("returned unexpected result: MMScore %v is bigger want bdminm.RMMScore %v is bigger", bdminm.RMMScore, bdminm.MMScore)
	}

	if bdminm.RMMScore == bdminm.MMScore {
		t.Errorf("returned unexpected result: MMScore %v is equal to bdminm.RMMScore.  want bdminm.rMMScore %v is bigger", bdminm.RMMScore, bdminm.MMScore)
	}
}

func TestBiDirectionalMinMatch_GetFrequency(t *testing.T) {
	bdminm.LoadDict()

	reality, err := bdminm.GetFrequency(textForBiDirectionalMatch)
	checkErr(err, t)

	result := strings.Split(expectedForBiDirectionalMinMatch, sep)
	expected := GetFrequency(result)

	if !reflect.DeepEqual(expected, reality) {
		t.Errorf(errorFormat,
			reality, expected)
	}
}
