package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{100, 4, 200, 1, 3, 2}, 4},
	{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
}

func TestLongestConsecutive(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if longestConsecutive(tcs[i].In) != tcs[i].Expect {
			t.Errorf("longest consecutive test failed on case: %d\n", i)
		}
	}
}
