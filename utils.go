package gotokenizer

import "errors"

func CheckDictIsLoaded(dict *Dict) error {
	if dict == nil {
		return errors.New("please load dictionary")
	}
	return nil
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func GetFrequency(result []string) map[string]int {
	resultMap := make(map[string]int)

	for _, v := range result {
		if _, ok := resultMap[v]; ok {
			resultMap[v]++
		} else {
			resultMap[v] = 1
		}
	}

	return resultMap
}
