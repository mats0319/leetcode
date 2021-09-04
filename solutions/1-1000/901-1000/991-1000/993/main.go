package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2 <= nodes.count <= 100
// 1 <= node.value <= 100
func isCousins(root *TreeNode, x int, y int) bool {
	xValue, xDepth := getDepth(root, x, root.Val, 0)
	yValue, yDepth := getDepth(root, y, root.Val, 0)

	return xValue != yValue && xDepth == yDepth
}

func getDepth(node *TreeNode, value, fatherValue, fatherDepth int) (int, int) {
	if node == nil {
		return 0, 0
	}

	if node.Val == value {
		return fatherValue, fatherDepth
	}

	fv, fd := getDepth(node.Left, value, node.Val, fatherDepth+1)
	if fv != 0 || fd != 0 {
		return fv, fd
	}
	fv, fd = getDepth(node.Right, value, node.Val, fatherDepth+1)
	if fv != 0 || fd != 0 {
		return fv, fd
	}

	return 0, 0
}
