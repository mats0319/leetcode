package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := inorderTraversal(root, make([]int, 0))
	sum := 0
	for i := 0; i < len(res); i++ {
		if low <= res[i] && res[i] <= high {
			sum += res[i]
		}
	}

	return sum
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
