package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"aacecaaa", "aaacecaaa"},
	{"abcd", "dcbabcd"},
}

func TestShortestPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if shortestPalindrome(tcs[i].In) != tcs[i].Expect {
			t.Errorf("shortest palindrome test failed on case: %d\n", i)
		}
	}
}
