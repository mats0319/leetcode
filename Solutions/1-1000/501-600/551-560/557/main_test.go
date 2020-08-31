package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
}

func TestReverseWords(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if reverseWords(tcs[i].In) != tcs[i].Expect {
			t.Errorf("reverse words test failed on case: %d\n", i)
		}
	}
}
