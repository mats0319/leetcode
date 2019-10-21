package rem

import "testing"

type TestCase struct {
	S string
	P string
}

var testCase = []TestCase{
	{"aa", "a"},
	{"aa", "a*"},
	{"ab", ".*"},
	{"aab", "c*a*b"},
	{"mississippi", "mis*is*p*."},
}

var result = []bool{
	false,
	true,
	true,
	true,
	false,
}

func TestIsMatch(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatch(tcs[i].S, tcs[i].P) != result[i] {
			t.Errorf("regular expression matching test failed on case: %d", i)
		}
	}
}

func TestIsMatchDP(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatchDP(tcs[i].S, tcs[i].P) != result[i] {
			t.Errorf("regular expression matching solution: 2 test failed on case: %d", i)
		}
	}
}
