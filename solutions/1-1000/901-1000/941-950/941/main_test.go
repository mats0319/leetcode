package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect bool
}{
	// test cases here
	{[]int{1,3,2}, true},
}

func TestValidMountainArray(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if validMountainArray(tcs[i].In) != tcs[i].Expect {
			t.Errorf("valid mountain array test failed on case: %d\n", i)
		}
	}
}
