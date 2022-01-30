package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect [][]int
}{
	// test cases here
	{[][]int{{0}, {0}, {0}, {0}, {0}}, [][]int{{0}, {0}, {0}, {0}, {0}}},
	{[][]int{{0}, {0}, {0}, {1}, {1}}, [][]int{{0}, {0}, {0}, {1}, {2}}},
	{[][]int{{0, 0, 0, 1, 1}}, [][]int{{0, 0, 0, 1, 2}}},
	{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
	{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
}

func TestUpdateMatrix(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnTwoDimensionalIntSliceStrict(updateMatrix(tcs[i].In), tcs[i].Expect) {
			t.Errorf("update matrix test failed on case: %d\n", i)
		}
	}
}

func compareOnTwoDimensionalIntSliceStrict(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := 0; i < len(a) && isEqual; i++ {
		if len(a[i]) != len(b[i]) {
			isEqual = false
			break
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				isEqual = false
				break
			}
		}
	}

	return isEqual
}
