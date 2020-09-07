package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect []int
}{
	// test cases here
	{},
}

func TestTopKFrequent(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(topKFrequent(tcs[i].In.IntSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("top k frequent test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[int]int)
	for i := range a {
		ma[a[i]]++
	}

	isEqual := true
	for i := range b {
		_, ok := ma[b[i]]
		if !ok {
			isEqual = false
			break
		}

		ma[b[i]]--
		if ma[b[i]] <= 0 {
			delete(ma, b[i])
		}
	}

	return isEqual
}
