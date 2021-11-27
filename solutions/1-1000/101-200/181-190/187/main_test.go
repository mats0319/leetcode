package mario

import "testing"

var testCase = []struct {
	In     string
	Expect []string
}{
	// test cases here
	{"AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnStringSlice(findRepeatedDnaSequences(tcs[i].In), tcs[i].Expect) {
			t.Errorf("find repeated dna sequences test failed on case: %d\n", i)
		}
	}
}

func compareOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[string]int, len(a))
	for i := range a {
		aMap[a[i]]++
		aMap[b[i]]--
	}

	isEqual := true
	for _, v := range aMap {
		if v != 0 {
			isEqual = false
			break
		}
	}

	return isEqual
}
