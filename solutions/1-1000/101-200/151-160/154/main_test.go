package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{2, 2, 2, 0, 1}, 0},
	{[]int{3, 1, 3}, 1},
}

func TestFindMin(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findMin(tcs[i].In) != tcs[i].Expect {
			t.Errorf("find min test failed on case: %d\n", i)
		}
	}
}
