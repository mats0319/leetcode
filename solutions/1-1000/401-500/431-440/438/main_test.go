package mario

import (
	"testing"
)

var testCase = []struct {
	In     []string
	Expect []int
}{
	// test cases here
	{[]string{"cbaebabacd", "abc"}, []int{0, 6}},
	{[]string{"abab", "ab"}, []int{0, 1, 2}},
}

func TestFindAnagrams(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(findAnagrams(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("find anagrams test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[int]int, 0) // value - times
	for i := range a {
		aMap[a[i]]++
		aMap[b[i]]--
	}

	isEqual := true
	for i := range aMap {
		if aMap[i] != 0 {
			isEqual = false
			break
		}
	}

	return isEqual
}
