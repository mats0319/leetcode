package mario

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// k <= nodes.count
func kthSmallest(root *TreeNode, k int) int {
	return inOrderTraversal(root, make([]int, 0))[k-1]
}

func inOrderTraversal(root *TreeNode, traversal []int) []int {
	if root == nil {
		return traversal
	}

	traversal = inOrderTraversal(root.Left, traversal)

	traversal = append(traversal, root.Val)

	traversal = inOrderTraversal(root.Right, traversal)

	return traversal
}
