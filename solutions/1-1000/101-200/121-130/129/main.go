package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return preorderTraversal(root, 0, 0)
}

func preorderTraversal(root *TreeNode, sum, curr int) int {
	if root == nil {
		return sum
	}

	curr = curr*10 + root.Val

	if root.Left == nil && root.Right == nil {
		sum += curr
	} else {
		sum = preorderTraversal(root.Left, sum, curr)

		sum = preorderTraversal(root.Right, sum, curr)
	}

	return sum
}
