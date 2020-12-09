package mario

import "testing"

type In struct {
	ByteSlice0 []byte
	Int0       int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2}, 8},
	{In{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0}, 6},
	{In{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2}, 16},
}

func TestLeastInterval(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if leastInterval(tcs[i].In.ByteSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("least interval test failed on case: %d\n", i)
		}
	}
}
