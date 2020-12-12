package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	// test cases here
	{"RD", "Radiant"},
	{"RDD", "Dire"},
	{"DRRDRDRDRDDRDRDR", "Radiant"},
}

func TestPredictPartyVictory(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if predictPartyVictory(tcs[i].In) != tcs[i].Expect {
			t.Errorf("predict party victory test failed on case: %d\n", i)
		}
	}
}
