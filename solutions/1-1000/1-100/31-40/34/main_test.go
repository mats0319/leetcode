package mario

import (
	"testing"
)

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect []int
}{
	// test cases here
	{In{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
	{In{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
	{In{[]int{2, 2}, 3}, []int{-1, -1}},
	{In{[]int{1}, 1}, []int{0, 0}},
	{In{[]int{2, 2}, 2}, []int{0, 1}},
}

func TestSearchRange(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(searchRange(tcs[i].In.IntSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("search range test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) (result bool) {
	if len(a) != len(b) {
		result = false
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
