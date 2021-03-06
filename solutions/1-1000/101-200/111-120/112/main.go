package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return targetSum == 0
	}

	return calcSum(root, 0, targetSum)
}

func calcSum(root *TreeNode, sum, target int) bool {
	if root.Left == nil && root.Right == nil {
		return sum+root.Val == target
	}

	isEqual := false

	if root.Left != nil {
		isEqual = isEqual || calcSum(root.Left, sum+root.Val, target)
	}

	if !isEqual && root.Right != nil {
		isEqual = isEqual || calcSum(root.Right, sum+root.Val, target)
	}

	return isEqual
}
