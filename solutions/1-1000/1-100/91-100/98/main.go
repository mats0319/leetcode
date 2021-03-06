package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	inOrderValue := inOrderTraversal(root, make([]int, 0))

	isValid := true
	for i := 0; i < len(inOrderValue)-1; i++ {
		if inOrderValue[i] >= inOrderValue[i+1] {
			isValid = false
			break
		}
	}

	return isValid
}

func inOrderTraversal(root *TreeNode, value []int) []int {
	if root == nil {
		return value
	}

	value = inOrderTraversal(root.Left, value)

	value = append(value, root.Val)

	value = inOrderTraversal(root.Right, value)

	return value
}
