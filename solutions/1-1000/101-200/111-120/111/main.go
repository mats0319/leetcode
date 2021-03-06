package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) (depth int) {
	if root == nil {
		return
	}

	layer := []*TreeNode{root}
	for len(layer) > 0 {
		depth++
		nextLayer := make([]*TreeNode, 0, len(layer)*2)

		for _, node := range layer {
			if node.Left == nil && node.Right == nil {
				nextLayer = nil // for exit outside loop
				break
			}

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
