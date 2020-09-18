package mario

import "testing"

type In struct {
	IstNodeSlice0 *ListNode
	Int0          int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{makeList(1,2,3,4,5), 4}, 2},
}

func TestKthToLast(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if kthToLast(tcs[i].In.IstNodeSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("kth to last test failed on case: %d\n", i)
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
