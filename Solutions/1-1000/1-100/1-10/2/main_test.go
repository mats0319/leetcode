package mario

import (
	"github.com/mats9693/leetcode/utils/compare"
	"testing"
)

var testCase = []struct {
	In     []*cmp.ListNode
	Expect *cmp.ListNode
}{
	{[]*cmp.ListNode{cmp.MakeList(2, 4, 3), cmp.MakeList(5, 6, 4)}, cmp.MakeList(7, 0, 8)},
}

func TestAddTwoNumbers(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnList(addTwoNumbers(tcs[i].In[0], tcs[i].In[1]), tcs[i].Expect) {
			t.Errorf("add two numbers test failed on case: %d", i)
		}
	}
}
