package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"bcabc", "abc"},
	{"cbacdcbc", "acdb"},
	{"bbcaac", "bac"},
}

func TestRemoveDuplicateLetters(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if removeDuplicateLetters(tcs[i].In) != tcs[i].Expect {
			t.Errorf("remove duplicate letters test failed on case: %d\n", i)
		}
	}
}
