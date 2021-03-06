package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	layer := []*TreeNode{root}
	fromLeft := true
	res := make([][]int, 0)
	for len(layer) > 0 {
		res = append(res, setValue(layer, fromLeft))

		nextLayer := make([]*TreeNode, 0, len(layer)*2)

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
		fromLeft = !fromLeft
	}

	return res
}

func setValue(slice []*TreeNode, fromLeft bool) []int {
	res := make([]int, len(slice))
	if fromLeft {
		for i := 0; i < len(slice); i++ {
			res[i] = slice[i].Val
		}
	} else {
		index := 0
		for i := len(slice) - 1; i >= 0; i-- {
			res[index] = slice[i].Val
			index++
		}
	}

	return res
}
