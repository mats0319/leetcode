package mario

import "testing"

var testCase = []struct {
	In *TreeNode
	Expect *TreeNode
}{
	// test cases here
	{In: makeTree([]*IntWithNil{
		{1, false},
		{3, false},
		{0, true},
		{0, true},
		{2, false},
	}), Expect: makeTree([]*IntWithNil{
		{3, false},
		{1, false},
		{0, true},
		{0, true},
		{2, false},
	})},
	{In: makeTree([]*IntWithNil{
		{3, false},
		{1, false},
		{4, false},
		{0, true},
		{0, true},
		{2, false},
	}), Expect: makeTree([]*IntWithNil{
		{2, false},
		{1, false},
		{4, false},
		{0, true},
		{0, true},
		{3, false},
	})},
}

func TestRecoverTree(t *testing.T) {
	//tcs := testCase
	//for i := range tcs {
	//	recoverTree(tcs[i].In)
	//	if !isSameTree(tcs[i].In, tcs[i].Expect) {
	//		t.Errorf("recover tree test failed on case: %d\n", i)
	//	}
	//}
}

type IntWithNil struct {
	Value int
	IsNil bool
}

func makeTree(slice []*IntWithNil) *TreeNode {
	if len(slice) <= 0 || (len(slice) == 1 && slice[0].IsNil) {
		return nil
	}

	root := &TreeNode{Val: slice[0].Value}
	nodeSlice := make([]*TreeNode, 0)
	nodeSlice = append(nodeSlice, root)

	leftFlag := true
	for i, j := 1, 0; i < len(slice) && j < len(nodeSlice); i++ {
		for ; j < len(nodeSlice) && nodeSlice[j] == nil; j++ {
		}

		var node *TreeNode // nil
		if !slice[i].IsNil {
			node = &TreeNode{Val: slice[i].Value}
		}

		if leftFlag {
			nodeSlice[j].Left = node
			leftFlag = !leftFlag
			nodeSlice = append(nodeSlice, nodeSlice[j].Left)
		} else {
			nodeSlice[j].Right = node
			leftFlag = !leftFlag
			nodeSlice = append(nodeSlice, nodeSlice[j].Right)
			j++
		}
	}

	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
