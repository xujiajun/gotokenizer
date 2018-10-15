package go_tokenizer

import "testing"

func TestStopTokens_IsStopToken(t *testing.T) {
	stopTokens := NewStopTokens()
	stopTokens.Load(stopTokensDictPath)
	testNotToken := "xxx"
	if stopTokens.IsStopToken(testNotToken) {
		t.Errorf(errorFormat,
			!stopTokens.IsStopToken(testNotToken), stopTokens.IsStopToken(testNotToken))
	}
	testIsToken := "."
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = " "
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "。"
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "【"
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "】"
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "["
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "]"
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = ","
	checkIsStopTokens(stopTokens, testIsToken, t)
	testIsToken = "的"
	checkIsStopTokens(stopTokens, testIsToken, t)
}

func checkIsStopTokens(stopTokens *StopTokens, testIsToken string, t *testing.T) {
	if !stopTokens.IsStopToken(testIsToken) {
		t.Errorf(errorFormat,
			!stopTokens.IsStopToken(testIsToken), stopTokens.IsStopToken(testIsToken))
	}
}
