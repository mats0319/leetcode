package mario

import (
	"testing"
)

var testCase = []struct {
	In     *TreeNode
	Expect int
}{
	// test cases here
	{makeTree([]*IntWithNil{
		{3, false},
		{2, false},
		{3, false},
		{0, true},
		{3, false},
		{0, true},
		{1, false},
	}), 7},
	{makeTree([]*IntWithNil{
		{3, false},
		{4, false},
		{5, false},
		{1, false},
		{3, false},
		{0, true},
		{1, false},
	}), 9},
	{makeTree([]*IntWithNil{
		{4, false},
		{1, false},
		{0, true},
		{2, false},
		{0, true},
		{3, false},
	}), 7},
}

func TestRob(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if rob(tcs[i].In) != tcs[i].Expect {
			t.Errorf("rob test failed on case: %d\n", i)
		}
	}
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
		if slice[i].IsNil {
			continue
		}

		if leftFlag {
			nodeSlice[j].Left = &TreeNode{Val: slice[i].Value}
			leftFlag = !leftFlag
			nodeSlice = append(nodeSlice, nodeSlice[j].Left)
		} else {
			nodeSlice[j].Right = &TreeNode{Val: slice[i].Value}
			leftFlag = !leftFlag
			nodeSlice = append(nodeSlice, nodeSlice[j].Right)
			j++
		}
	}

	return root
}
