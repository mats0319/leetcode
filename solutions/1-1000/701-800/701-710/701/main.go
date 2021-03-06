package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	pre := root
	for p := root; p != nil; {
		pre = p

		if val > p.Val {
			p = p.Right
		} else { // if val <= root.Val
			p = p.Left
		}
	}

	if val > pre.Val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}

	return root
}
