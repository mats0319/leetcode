package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"owoztneoer", "012"},
	{"zeroonetwothreefourfivesixseveneightnine", "0123456789"},
}

func TestOriginalDigits(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if originalDigits(tcs[i].In) != tcs[i].Expect {
			t.Errorf("original digits test failed on case: %d\n", i)
		}
	}
}
