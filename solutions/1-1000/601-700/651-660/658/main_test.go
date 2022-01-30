package mario

import "testing"

type In struct {
	Int0      int
	Int1      int
	IntSlice0 []int
}

var testCase = []struct {
	In     In
	Expect []int
}{
	// test cases here
	{In{30, 1, []int{1, 10, 15, 25, 35, 45, 50, 59}}, []int{25}},
}

func TestFindClosestElements(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(findClosestElements(tcs[i].In.IntSlice0, tcs[i].In.Int1, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("find closest elements test failed on case: %d\n", i)
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
