package mario

import "testing"

var testCase = []struct {
	In     string
	Expect int
}{
	// test cases here
	{"abc", 3},
	{"aaa", 6},
}

func TestCountSubstrings(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if countSubstrings(tcs[i].In) != tcs[i].Expect {
			t.Errorf("count substrings test failed on case: %d\n", i)
		}
	}
}
