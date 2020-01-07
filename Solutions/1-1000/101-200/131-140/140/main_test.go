package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []string
	Expect []string
}{
	{[]string{"catsanddog", "cat", "cats", "and", "sand", "dog"},
		[]string{"cats and dog", "cat sand dog"}},
	{[]string{"pineapplepenapple", "apple", "pen", "applepen", "pine", "pineapple"},
		[]string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
	{[]string{"catsandog", "cats", "dog", "sand", "and", "cat"},
		[]string{}},
}

func TestWordBreak(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnStringSlice(wordBreak(tcs[i].In[0], tcs[i].In[1:]), tcs[i].Expect) {
			t.Errorf("word break test failed on case: %d", i)
		}
	}
}
