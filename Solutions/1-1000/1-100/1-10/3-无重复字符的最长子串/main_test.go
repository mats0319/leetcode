package mario

import (
	"testing"
)

var testCase = []struct {
	In     string
	Expect int
}{
	{"abcabcbba", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{" ", 1},
	{"aab", 2},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if lengthOfLongestSubstring(tcs[i].In) != tcs[i].Expect {
			t.Errorf("length of longest substring test failed on case: %d", i)
		}
	}
}
