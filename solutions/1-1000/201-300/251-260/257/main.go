package mario

import "strconv"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// binaryTreePaths root != nil
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)

    return dfs(root, "", res)
}

func dfs(root *TreeNode, str string, res []string) []string {
    str += "->"+strconv.Itoa(root.Val)

    if root.Left == nil && root.Right == nil {
        res = append(res, str[2:]) // skip first '->'
        return res
    }

    if root.Left != nil {
        res = dfs(root.Left, str, res)
    }
    if root.Right != nil {
        res = dfs(root.Right, str, res)
    }

    return res
}
