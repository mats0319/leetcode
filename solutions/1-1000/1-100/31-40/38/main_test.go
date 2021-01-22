package mario

import "testing"

var testCase = []struct {
	In     int
	Expect string
}{
	// test cases here
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
	{5, "111221"},
}

func TestCountAndSay(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if countAndSay(tcs[i].In) != tcs[i].Expect {
			t.Errorf("count and say test failed on case: %d\n", i)
		}
	}
}
