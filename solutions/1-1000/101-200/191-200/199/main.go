package mario

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		nextList := make([]*TreeNode, 0, len(nodeList)*2)

		rightValue := 0
		for len(nodeList) > 0 {
			node := nodeList[0]
			nodeList = nodeList[1:]

			rightValue = node.Val

			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}
			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}

		res = append(res, rightValue)

		nodeList = nextList
	}

	return res
}
