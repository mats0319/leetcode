package mario

import "testing"

type In struct {
	IstNodeSlice0 *ListNode
	Int0          int
}

var testCase = []struct {
	In     In
	Expect *ListNode
}{
	// test cases here
	{In{makeList(1,2,3,4,5), 2}, makeList(4,5)},
}

func TestGetKthFromEnd(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(getKthFromEnd(tcs[i].In.IstNodeSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("get kth from end test failed on case: %d\n", i)
		}
	}
}

func makeList(value ...int) *ListNode {
	pre := &ListNode{}
	p := pre
	for i := range value {
		p.Next = &ListNode{Val: value[i]}
		p = p.Next
	}

	return pre.Next
}

func compareOnList(a, b *ListNode) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	aCopy, bCopy := a, b
	for aCopy != nil && bCopy != nil && aCopy.Val == bCopy.Val {
		aCopy = aCopy.Next
		bCopy = bCopy.Next
	}

	return aCopy == nil && bCopy == nil
}
