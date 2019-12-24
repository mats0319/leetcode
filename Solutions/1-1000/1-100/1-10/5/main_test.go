package mario

import (
	"testing"
)

var testCase = []struct {
	In     string
	Expect string
}{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"ac", "a"},
}

func TestLongestPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestPalindrome(tcs[i].In) != tcs[i].Expect {
			t.Errorf("longest palindromic substring test failed on case: %d", i)
		}
	}
}
