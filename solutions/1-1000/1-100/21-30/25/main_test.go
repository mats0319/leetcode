package mario

import "testing"

type In struct {
	ListNodeSlice0 *ListNode
	Int0           int
}

var testCase = []struct {
	In     In
	Expect *ListNode
}{
	// test cases here
	{In{makeList(1, 2), 2}, makeList(2, 1)},
	{In{makeList(1, 2, 3, 4, 5), 3}, makeList(3, 2, 1, 4, 5)},
	{In{makeList(1, 2, 3, 4, 5), 2}, makeList(2, 1, 4, 3, 5)},
	{In{makeList(1, 2, 3, 4, 5), 1}, makeList(1, 2, 3, 4, 5)},
	{In{makeList(1), 1}, makeList(1)},
}

func TestReverseKGroup(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnList(reverseKGroup(tcs[i].In.ListNodeSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("reverse k group test failed on case: %d\n", i)
		}
	}
}

func makeList(values ...int) *ListNode {
	pre := &ListNode{}
	p := pre
	for i := range values {
		p.Next = &ListNode{Val: values[i]}
		p = p.Next
	}

	return pre.Next
}

func compareOnList(a, b *ListNode) bool {
	isEqual := true
	for a != nil && b != nil {
		if a.Val != b.Val {
			isEqual = false
			break
		}

		a = a.Next
		b = b.Next
	}

	return isEqual && a == b
}
