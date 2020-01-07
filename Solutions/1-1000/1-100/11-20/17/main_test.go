package mario

import (
	"github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     string
	Expect []string
}{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestLetterCombinations(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnStringSlice(letterCombinations(tcs[i].In), tcs[i].Expect) {
			t.Errorf("letter combinations test failed on case: %d", i)
		}
	}
}
