package mario

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	layer := []*Node{root}
	for len(layer) > 0 {
		nextLayer := make([]*Node, 0)
		valueLayer := make([]int, 0)

		for len(layer) > 0 {
			node := layer[0]
			layer = layer[1:]

			valueLayer = append(valueLayer, node.Val)
			nextLayer = append(nextLayer, node.Children...)
		}

		res = append(res, valueLayer)
		layer = nextLayer
	}

	return res
}
