package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MaxValue struct {
	PickRoot    int
	NotPickRoot int
}

func rob(root *TreeNode) int {
	// 左子树、右子树、根节点，递归
	if isNil(root) {
		return 0
	}
	if isLeafNode(root) {
		return root.Val
	}

	value := robTree(root)

	return big(value.PickRoot, value.NotPickRoot)
}

func robTree(root *TreeNode) *MaxValue {
	if isNil(root) {
		return &MaxValue{} // testdata 0 value
	}
	if isLeafNode(root) {
		value := root.Val
		if value < 0 {
			value = 0
		}
		return &MaxValue{PickRoot: value, NotPickRoot: 0}
	}

	leftValue := robTree(root.Left)
	rightValue := robTree(root.Right)

	value := &MaxValue{
		PickRoot:    leftValue.NotPickRoot + rightValue.NotPickRoot,
		NotPickRoot: leftValue.PickRoot + rightValue.PickRoot,
	}

	if root.Val > 0 {
		value.PickRoot += root.Val
	}
	if value.PickRoot < value.NotPickRoot {
		value.PickRoot = value.NotPickRoot
	}

	return value
}

func isNil(node *TreeNode) bool {
	return node == nil
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func big(a, b int) (bigger int) {
	if a > b {
		bigger = a
	} else {
		bigger = b
	}

	return
}
