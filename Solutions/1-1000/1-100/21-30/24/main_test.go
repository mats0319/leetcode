package mario

import (
	cmp "github.com/mats9693/utils/compare"
	"testing"
)

var testCase = []struct {
	In     *cmp.ListNode
	Expect *cmp.ListNode
}{
	// test cases here
	{cmp.MakeList(1, 2, 3, 4), cmp.MakeList(2, 1, 4, 3)},
	{cmp.MakeList(1, 2, 3), cmp.MakeList(2, 1, 3)},
}

func TestSwapPairs(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		result := swapPairs(tcs[i].In)
		if !cmp.CompareOnList(result, tcs[i].Expect) {
			cmp.PrintList(result)
			t.Errorf("swap pairs test failed on case: %d\n", i)
		}
	}
}
