package mario

import "testing"

var testCase = []struct {
	In     string
	Expect int
}{
	// test cases here
	{"abcabcbba", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{" ", 1},
	{"aab", 2},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for i := range testCase {
		if lengthOfLongestSubstring(testCase[i].In) != testCase[i].Expect {
			t.Logf("length of longest substring test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
