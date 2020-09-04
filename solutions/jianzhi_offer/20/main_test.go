package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	{"G76", false},
	{"1", true},
	{"1 ", true},
	{"3. ", true},
	{"+100", true},
	{"5e2", true},
	{"-123", true},
	{"3.1416", true},
	{"-1E-16", true},
	{"0123", true},
	{"12e", false},
	{"1a3.14", false},
	{"1.2.3", false},
	{"+-5", false},
	{"12e+5.4", false},
}

func TestIsNumber(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isNumber(tcs[i].In) != tcs[i].Expect {
			t.Errorf("is number test failed on case: %d\n", i)
		}
	}
}
