package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect string
}{
	// test cases here
	{[]int{1, 2}, "0.5"},
	{[]int{2, 1}, "2"},
	{[]int{2, 3}, "0.(6)"},
	{[]int{4, 333}, "0.(012)"},
	{[]int{1, 5}, "0.2"},
	{[]int{22, 7}, "3.(142857)"},
	{[]int{-50, 8}, "-6.25"},
	{[]int{7, -12}, "-0.58(3)"},
}

func TestFractionToDecimal(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if fractionToDecimal(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("fraction to decimal test failed on case: %d\n", i)
		}
	}
}
