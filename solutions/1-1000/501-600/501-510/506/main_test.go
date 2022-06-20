package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect []string
}{
	// test cases here
	{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
}

func TestFindRelativeRanks(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnStringSlice(findRelativeRanks(tcs[i].In), tcs[i].Expect) {
			t.Errorf("find relative ranks test failed on case: %d\n", i)
		}
	}
}

func compareOnStringSlice(a, b []string) bool {
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
