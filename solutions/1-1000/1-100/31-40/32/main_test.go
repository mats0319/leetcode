package mario

import "testing"

var testCase = []struct {
	In     string
	Expect int
}{
	{"(()", 2},
	{")()())", 4},
	{"()", 2},
	{"())", 2},
	{"()(())", 6},
	{"()(()", 2},
}

func TestLongestValidParentheses(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestValidParentheses(tcs[i].In) != tcs[i].Expect {
			t.Errorf("longest valid parentheses test failed on case: %d\n", i)
		}
	}
}

func TestLongestValidParenthesesWithConstantStorage(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestValidParenthesesWithConstantStorage(tcs[i].In) != tcs[i].Expect {
			t.Errorf("longest valid parentheses with constant storage test failed on case: %d\n", i)
		}
	}
}
