package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect []int
}{
	// test cases here
	{In: []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41},
		Expect: []int{10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0}},
	{In: []int{2, 0, 1}, Expect: []int{2, 0, 0}},
	{In: []int{5, 2, 6, 1}, Expect: []int{2, 1, 1, 0}},
}

func TestCountSmaller(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(countSmaller(tcs[i].In), tcs[i].Expect) {
			t.Errorf("count smaller test failed on case: %d\n", i)
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
