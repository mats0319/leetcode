package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val > root.Val {
			root = root.Right
		} else if val < root.Val {
			root = root.Left
		} else { // if val == root.Val
			break
		}
	}

	return root
}
