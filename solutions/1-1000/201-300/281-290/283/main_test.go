package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect []int
}{
	// test cases here
	{In: []int{1, 0, 1}, Expect: []int{1, 1, 0}},
	{In: []int{0}, Expect: []int{0}},
	{In: []int{0, 1, 0, 3, 12}, Expect: []int{1, 3, 12, 0, 0}},
}

func TestMoveZeroes(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		moveZeroes(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("move zeroes test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
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
