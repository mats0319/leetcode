package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{2, 1, 4, 7, 3, 2, 5}, 5},
	{[]int{1, 4, 7, 3, 2}, 5},
}

func TestLongestMountain(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestMountain(tcs[i].In) != tcs[i].Expect {
			t.Errorf("longest mountain test failed on case: %d\n", i)
		}
	}
}
