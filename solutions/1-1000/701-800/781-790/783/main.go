package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	traversal := inorderTraversal(root, make([]int, 0))

	min := traversal[1] - traversal[0]
	for i := 2; i < len(traversal); i++ {
		if traversal[i]-traversal[i-1] < min {
			min = traversal[i] - traversal[i-1]
		}
	}

	return min
}

func inorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = inorderTraversal(root.Left, res)

	res = append(res, root.Val)

	res = inorderTraversal(root.Right, res)

	return res
}
