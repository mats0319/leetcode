package mario

type Node struct {
    Val int
    Children []*Node
}

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    return dfs(root)
}

func dfs(root *Node) int {
    max := 0
    for _, v := range root.Children {
        depth := dfs(v)
        if max < depth {
            max = depth
        }
    }

    return max+1
}
