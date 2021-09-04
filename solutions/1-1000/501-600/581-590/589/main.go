package mario

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	for i := range root.Children {
		sequence := preorder(root.Children[i])
		if len(sequence) > 0 {
			res = append(res, sequence...)
		}
	}

	return res
}
