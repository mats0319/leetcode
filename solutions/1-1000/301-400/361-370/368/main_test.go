package mario

import (
	"testing"
)

var testCase = []struct {
	In     []int
	Expect []int
}{
	// test cases here
	//{[]int{3, 4, 16, 8}, []int{16, 8, 4}},
	{[]int{4, 8, 10, 240}, []int{240, 8, 4}},
}

func TestLargestDivisibleSubset(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(largestDivisibleSubset(tcs[i].In), tcs[i].Expect) {
			t.Errorf("largest divisible subset test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
