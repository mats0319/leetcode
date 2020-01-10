package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
}

func TestMaxSubArray(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxSubArray(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max sub array test failed on case: %d\n", i)
		}
	}
}
