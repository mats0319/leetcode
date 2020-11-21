package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	{"1,#,#,1", false},
	{"1,#", false},
}

func TestIsValidSerialization(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isValidSerialization(tcs[i].In) != tcs[i].Expect {
			t.Errorf("is valid serialization test failed on case: %d\n", i)
		}
	}
}
