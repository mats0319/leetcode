package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	//{"abab", true},
	//{"aba", false},
	//{"abcabcabcabc", true},
	//{"ababba", false},
	{"babbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbbbabbaaabbb", true},
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if repeatedSubstringPattern(tcs[i].In) != tcs[i].Expect {
			t.Errorf("repeated substring pattern test failed on case: %d\n", i)
		}
	}
}
