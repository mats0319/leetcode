package main

import "testing"

var testCase = []struct{
	In string
	Except []string
}{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestLetterCombinations(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnStringSlice(letterCombinations(tcs[i].In), tcs[i].Except) {
			t.Errorf("letter combinations test failed on case: %d", i)
		}
	}
}

func compareOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	var bMap = make(map[string]int, len(b))
	for i := range b {
		bMap[b[i]] = 0
	}

	for i := range a {
		if _, ok := bMap[a[i]]; ok {
			delete(bMap, a[i])
		} else {
			return false
		}
	}

	return len(bMap) == 0
}
