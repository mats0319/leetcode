package mario

import "testing"

var testCase = []struct {
	In     int
	Expect string
}{
	{3, "III"},
	{4, "IV"},
	{9, "IX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
}

func TestIntToRoman(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if intToRoman(tcs[i].In) != tcs[i].Expect {
			t.Errorf("int to roman test failed on case: %d\n", i)
		}
	}
}
