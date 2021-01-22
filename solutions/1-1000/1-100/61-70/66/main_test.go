package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect []int
}{
	// test cases here
	{[]int{1,2,3}, []int{1,2,4}},
	{[]int{9}, []int{1,0}},
}

func TestPlusOne(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(plusOne(tcs[i].In), tcs[i].Expect) {
			t.Errorf("plus one test failed on case: %d\n", i)
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
