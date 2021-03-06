package mario

import "testing"

var testCase = []struct {
	In     string
	Expect []int
}{
	// test cases here
	{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
}

func TestPartitionLabels(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(partitionLabels(tcs[i].In), tcs[i].Expect) {
			t.Errorf("partition labels test failed on case: %d\n", i)
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
