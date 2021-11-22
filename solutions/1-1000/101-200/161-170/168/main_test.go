package mario

import "testing"

var testCase = []struct {
	In     int
	Expect string
}{
	// test cases here
	{1, "A"},
	{701, "ZY"},
}

func TestConvertToTitle(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if convertToTitle(tcs[i].In) != tcs[i].Expect {
			t.Errorf("convert to title test failed on case: %d\n", i)
		}
	}
}
