package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect int
}{
	// test cases here
	{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
	{[][]int{{30, 50}, {12, 2}, {3, 4}, {12, 15}}, 3},
}

func TestMaxEnvelopes(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxEnvelopes(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max envelopes test failed on case: %d\n", i)
		}
	}
}
