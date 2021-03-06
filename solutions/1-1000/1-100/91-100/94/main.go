package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return traversal(root, make([]int, 0))
}

func traversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = traversal(root.Left, res)

	res = append(res, root.Val)

	res = traversal(root.Right, res)

	return res
}
