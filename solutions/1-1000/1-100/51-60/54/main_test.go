package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect []int
}{
	// test cases here
	{[][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
}

func TestSpiralOrder(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSliceStrict(spiralOrder(tcs[i].In), tcs[i].Expect) {
			t.Errorf("spiral order test failed on case: %d\n", i)
		}
	}
}

// strict: both value and order is equal
func compareOnIntSliceStrict(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
