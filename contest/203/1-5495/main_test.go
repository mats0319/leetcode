package mario

import (
    "testing"
)

type In struct {
	Int0      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect []int
}{
	// test cases here
	{In{4, []int{1,3,1,2}}, []int{1,2}},
	{In{2, []int{2,1,2,1,2,1,2,1,2}}, []int{2}},
	{In{7, []int{1,3,5,7}}, []int{1,2,3,4,5,6,7}},
	{In{3, []int{3,2,1,2,1,3,2,1,2,1,3,2,3,1}}, []int{1,3}},
	{In{6, []int{4,5}}, []int{4,5}},
}

func TestMostVisited(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(mostVisited(tcs[i].In.Int0, tcs[i].In.IntSlice0), tcs[i].Expect) {
			t.Errorf("most visited test failed on case: %d\n", i)
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
