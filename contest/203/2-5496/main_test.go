package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	// test cases here
	{[]int{2, 4, 1, 2, 7, 8}, 9},
	{[]int{2, 4, 5}, 4},
	{[]int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
}

func TestMaxCoins(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxCoins(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max coins test failed on case: %d\n", i)
		}
	}
}
