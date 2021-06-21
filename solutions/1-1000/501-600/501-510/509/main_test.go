package mario

import "testing"

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{3, 2},
	{4, 3},
}

func TestFib(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if fib(tcs[i].In) != tcs[i].Expect {
			t.Errorf("fib test failed on case: %d\n", i)
		}
	}
}
