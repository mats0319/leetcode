package lps

import (
	"testing"
)

var testCase = []string{
	"babad",
	"cbbd",
	"ac",
}

var result = []string{
	"bab",
	"bb",
	"a",
}

func TestLongestPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestPalindrome(tcs[i]) != result[i] {
			t.Errorf("longest palindromic substring test failed on case: %d", i)
		}
	}
}
