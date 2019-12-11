package mario

import (
    cmp "github.com/mats9693/leetcode/utils/compare"
    "testing"
)

var testCase = []struct {
    In []*cmp.ListNode
    Except *cmp.ListNode
}{
    {[]*cmp.ListNode{cmp.MakeList(1,2,4), cmp.MakeList(1,3,4)}, cmp.MakeList(1,1,2,3,4,4)},
    {[]*cmp.ListNode{cmp.MakeList(1), cmp.MakeList(1)}, cmp.MakeList(1,1)},
}

func TestMergeTwoLists(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if !cmp.CompareOnList(mergeTwoLists(tcs[i].In[0], tcs[i].In[1]), tcs[i].Except) {
            t.Errorf("merge two lists test failed on case: %d", i)
        }
    }
}
