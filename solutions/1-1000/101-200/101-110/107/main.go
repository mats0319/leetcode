package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	layerTraversal := make([][]int, 0)

	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		nextLayerNodeList := make([]*TreeNode, 0, 2*len(nodeList))

		valueList := make([]int, 0, len(nodeList))
		for len(nodeList) > 0 {
			node := nodeList[0]
			nodeList = nodeList[1:]

			valueList = append(valueList, node.Val)

			if node.Left != nil {
				nextLayerNodeList = append(nextLayerNodeList, node.Left)
			}
			if node.Right != nil {
				nextLayerNodeList = append(nextLayerNodeList, node.Right)
			}
		}

		layerTraversal = append(layerTraversal, valueList)

		nodeList = nextLayerNodeList
	}

	res := make([][]int, 0, len(layerTraversal))
	for i := len(layerTraversal) - 1; i >= 0; i-- {
		res = append(res, layerTraversal[i])
	}

	return res
}
