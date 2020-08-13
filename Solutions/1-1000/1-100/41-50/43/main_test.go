package mario

import "testing"

var testCase = []struct {
	In     []string
	Expect string
}{
	// test cases here
	{[]string{"2", "3"}, "6"},
	{[]string{"123", "456"}, "56088"},
	{[]string{"999", "999"}, "998001"},
	{[]string{"123456789", "987654321"}, "121932631112635269"},

}

func TestMultiply(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if multiply(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("multiply test failed on case: %d\n", i)
		}
	}
}
