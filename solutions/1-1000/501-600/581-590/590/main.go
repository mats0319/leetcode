package mario

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	return dfs(root)
}

func dfs(root *Node) []int {
	if len(root.Children) < 1 {
		return []int{root.Val}
	}

	res := make([]int, 0)
	for _, v := range root.Children {
		res = append(res, dfs(v)...)
	}

	res = append(res, root.Val)

	return res
}
