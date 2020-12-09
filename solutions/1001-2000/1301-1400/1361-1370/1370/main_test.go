package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"aaabbbcccddd", "abcddcbaabcd"},
}

func TestSortString(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if sortString(tcs[i].In) != tcs[i].Expect {
			t.Errorf("sort string test failed on case: %d\n", i)
		}
	}
}
