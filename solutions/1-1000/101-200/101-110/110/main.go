package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return postorderTraversal(root) >= 0
}

func postorderTraversal(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := postorderTraversal(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := postorderTraversal(root.Right)

	if rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
