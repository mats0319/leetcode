package mario

import "testing"

var testCase = []struct {
	A []int
	m int
	B []int
	n int
}{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 4, 5}, 3},
}

func TestMerge(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		merge(tcs[i].A, tcs[i].m, tcs[i].B, tcs[i].n)
		if !compareOnIntSlice(tcs[i].A, []int{1, 2, 2, 3, 4, 5}) {
			t.Errorf("merge test failed on case: %d\n", i)
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
