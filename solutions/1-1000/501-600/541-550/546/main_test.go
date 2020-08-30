package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{},
}

func TestRemoveBoxes(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if removeBoxes(tcs[i].In) != tcs[i].Expect {
			t.Errorf("remove boxes test failed on case: %d\n", i)
		}
	}
}
