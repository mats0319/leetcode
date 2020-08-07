package mario

import "testing"

var testCase = []struct {
	In     []*TreeNode
	Expect bool
}{
	//test cases here
	{In: []*TreeNode{
		makeTree([]*IntWithNil{
			{1, false},
			{2, false},
			{3, false},
		}),
		makeTree([]*IntWithNil{
			{1, false},
			{2, false},
			{3, false},
		}),
	}, Expect: true},
	{In: []*TreeNode{
		makeTree([]*IntWithNil{
			{1, false},
			{2, false},
		}),
		makeTree([]*IntWithNil{
			{1, false},
			{0, true},
			{2, false},
		}),
	}, Expect: false},
	{In: []*TreeNode{
		makeTree([]*IntWithNil{
			{1, false},
			{2, false},
			{1, false},
		}),
		makeTree([]*IntWithNil{
			{1, false},
			{1, false},
			{2, false},
		}),
	}, Expect: false},
}

func TestIsSameTree(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isSameTree(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
			t.Errorf("is same tree test failed on case: %d\n", i)
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
			leftFlag = !leftFlag
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
