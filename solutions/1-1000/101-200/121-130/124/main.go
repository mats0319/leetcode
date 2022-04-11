package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// loop each node, make it as 'root', which means all road is no matter with its parent node or brother nodes
// maintain max sum in recurse when calc from bottom to top
// for each node:
//   1. if node is a leaf, if node.val <= 0, do nothing; if node.val > 0, add node and calc node.max = node.val
//   2. if node is not a leaf, node.max = node.val + [all positive children.max]
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return big(recurse(root))
}

// recurse return 'root value' and 'max'
func recurse(root *TreeNode) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}

	rootSum := root.Val

	rootMax := root.Val

	max := rootSum

	left := 0
	if root.Left != nil {
		var leftMax int
		left, leftMax = recurse(root.Left)
		if left > 0 {
			rootMax += left
		}
		if max < leftMax {
			max = leftMax
		}
	}

	right := 0
	if root.Right != nil {
		var rightMax int
		right, rightMax = recurse(root.Right)
		if right > 0 {
			rootMax += right
		}
		if max < rightMax {
			max = rightMax
		}
	}

	if rootMax > rootSum { // left > 0 or right > 0
		rootSum += big(left, right)
	}

	return rootSum, big(rootMax, max) // root max >= root sum, always
}

func big(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
