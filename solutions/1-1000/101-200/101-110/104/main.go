package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	layer := make([]*TreeNode, 0)
	layer = append(layer, root)

	depth := 0
	for len(layer) > 0 {
		depth++
		nextLayer := make([]*TreeNode, 0, 2*len(layer))
		for _, node := range layer {
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		layer = nextLayer
	}

	return depth
}
