package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LayerNode struct {
	*TreeNode
	IsNil bool
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	layer := make([]*LayerNode, 0)
	if root.Left == nil {
		layer = append(layer, &LayerNode{IsNil: true})
	} else {
		layer = append(layer, &LayerNode{TreeNode: root.Left})
	}
	if root.Right == nil {
		layer = append(layer, &LayerNode{IsNil: true})
	} else {
		layer = append(layer, &LayerNode{TreeNode: root.Right})
	}

	isValid := true
ALL:
	for len(layer) > 0 {
		nextLayer := make([]*LayerNode, 0, 2*len(layer))

		for _, node := range layer {
			if node.IsNil {
				continue
			}

			if node.Left == nil {
				nextLayer = append(nextLayer, &LayerNode{IsNil: true})
			} else {
				nextLayer = append(nextLayer, &LayerNode{TreeNode: node.Left})
			}
			if node.Right == nil {
				nextLayer = append(nextLayer, &LayerNode{IsNil: true})
			} else {
				nextLayer = append(nextLayer, &LayerNode{TreeNode: node.Right})
			}
		}

		for left, right := 0, len(layer)-1; left < right; left, right = left+1, right-1 {
			if layer[left].IsNil && layer[right].IsNil {
				continue
			} else if (layer[left].IsNil || layer[right].IsNil) || layer[left].Val != layer[right].Val {
				isValid = false
				break ALL
			}
		}

		layer = nextLayer
	}

	return isValid
}
