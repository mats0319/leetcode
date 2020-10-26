package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []string
	Expect [][]int
}{
	// test cases here
	{In: []string{"abcd", "dcba", "lls", "s", "sssll"}, Expect: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
	{In: []string{"bat", "tab", "cat"}, Expect: [][]int{{0, 1}, {1, 0}}},
}

func TestPalindromePairs(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnTwoDimensionalSlice(palindromePairs(tcs[i].In), tcs[i].Expect) {
			t.Errorf("palindrome pairs test failed on case: %d\n", i)
		}
	}
}
