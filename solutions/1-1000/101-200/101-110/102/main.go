package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	layer := []*TreeNode{root}
	for len(layer) > 0 {
		layerValue := make([]int, 0, len(layer))
		newLayer := make([]*TreeNode, 0)
		for _, node := range layer {
			layerValue = append(layerValue, node.Val)

			if node.Left != nil {
				newLayer = append(newLayer, node.Left)
			}
			if node.Right != nil {
				newLayer = append(newLayer, node.Right)
			}
		}

		res = append(res, layerValue)
		layer = newLayer
	}

	return res
}
