package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return recurseTraversal(root, make([]int, 0))
}

func recurseTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	res = recurseTraversal(root.Left, res)

	res = recurseTraversal(root.Right, res)

	return res
}
