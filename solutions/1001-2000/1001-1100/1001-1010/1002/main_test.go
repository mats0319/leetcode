package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []string
	Expect []string
}{
	// test cases here
	{In: []string{"bella", "label", "roller"}, Expect: []string{"e", "l", "l"}},
	{In: []string{"cool", "lock", "cook"}, Expect: []string{"c", "o"}},
}

func TestCommonChars(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnStringSlice(commonChars(tcs[i].In), tcs[i].Expect) {
			t.Errorf("common chars test failed on case: %d\n", i)
		}
	}
}
