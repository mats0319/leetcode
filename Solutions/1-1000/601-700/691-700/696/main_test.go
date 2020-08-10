package mario

import "testing"

var testCase = []struct {
	In     string
	Expect int
}{
	// test cases here
	{"00110011", 6},
	{"10101", 4},
}

func TestCountBinarySubstrings(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if countBinarySubstrings(tcs[i].In) != tcs[i].Expect {
			t.Errorf("count binary substrings test failed on case: %d\n", i)
		}
	}
}
