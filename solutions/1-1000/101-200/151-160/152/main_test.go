package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{2, 3, -2, 4}, 6},
	{[]int{-2, 0, -1}, 0},
}

func TestMaxProduct(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if maxProduct(tcs[i].In) != tcs[i].Expect {
			t.Errorf("max product test failed on case: %d\n", i)
		}
	}
}
