package main

import "testing"

type In struct {
	S string
	P string
}

var testCase = []struct {
	In     In
	Expect bool
}{
	{In{"aa", "a"}, false},
	{In{"aa", "a*"}, true},
	{In{"ab", ".*"}, true},
	{In{"aab", "c*a*b"}, true},
	{In{"mississippi", "mis*is*p*."}, false},
}

func TestIsMatch(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatch(tcs[i].In.S, tcs[i].In.P) != tcs[i].Expect {
			t.Errorf("regular expression matching test failed on case: %d", i)
		}
	}
}

func TestIsMatchDP(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isMatchDP(tcs[i].In.S, tcs[i].In.P) != tcs[i].Expect {
			t.Errorf("regular expression matching solution: 2 test failed on case: %d", i)
		}
	}
}
