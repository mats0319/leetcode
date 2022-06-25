package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findBottomLeftValue root not nil
func findBottomLeftValue(root *TreeNode) (res int) {
	layer := []*TreeNode{root}
	for len(layer) > 0 {
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		res = layer[0].Val

		for len(layer) > 0 {
			node := layer[0]
			layer = layer[1:]

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}

			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		layer = nextLayer
	}

	return
}
