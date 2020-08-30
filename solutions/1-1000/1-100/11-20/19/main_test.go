package mario

import (
	"github.com/mats9693/utils/compare"
	"testing"
)

type In struct {
	Node *cmp.ListNode
	N    int
}

var testCase = []struct {
	In     In
	Expect *cmp.ListNode
}{
	{In{cmp.MakeList(1, 2, 3, 4, 5), 2}, cmp.MakeList(1, 2, 3, 5)},
	{In{cmp.MakeList(1), 1}, nil},
	{In{cmp.MakeList(1, 2), 2}, cmp.MakeList(2)},
}

func TestRemoveNthFromEnd(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !cmp.CompareOnList(removeNthFromEnd(tcs[i].In.Node, tcs[i].In.N), tcs[i].Expect) {
			t.Errorf("remove nth from end test failed on case: %d\n", i)
		}
	}
}
