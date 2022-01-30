package mario

import "testing"

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{10, 6},
}

func TestGuess(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		target = tcs[i].Expect
		if guessNumber(tcs[i].In) != tcs[i].Expect {
			t.Errorf("guess test failed on case: %d\n", i)
		}
	}
}
