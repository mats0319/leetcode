package mario

import "testing"

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{},
}

func TestF(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if f(tcs[i].In) != tcs[i].Expect {
			t.Errorf("f test failed on case: %d\n", i)
		}
	}
}
