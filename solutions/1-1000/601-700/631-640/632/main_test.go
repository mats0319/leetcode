package mario

import "testing"

var testCase = []struct {
	In     [][]int
	Expect []int
}{
	// test cases here
	{},
}

func TestSmallestRange(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(smallestRange(tcs[i].In), tcs[i].Expect) {
			t.Errorf("smallest range test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) (result bool) {
	if len(a) != len(b) {
		return
	}

	result = true
	for i := range a {
		if a[i] != b[i] {
			result = false
			break
		}
	}

	return
}
