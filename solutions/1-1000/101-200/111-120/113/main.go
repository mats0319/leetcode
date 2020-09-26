package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
    res = make([][]int, 0) // init every time, for lc

    dfs(root, sum, make([]int, 1))

    return res
}

func dfs(root *TreeNode, target int, solution []int) {
    if root == nil {
        return
    }

    target -= root.Val
    solution = append(solution, root.Val)

    if target == 0 && isLeafNode(root) {
        res = append(res, deepCopy(solution[1:]))
    }

    dfs(root.Left, target, solution)
    dfs(root.Right, target, solution)

    solution = solution[1:]

    return
}

func isLeafNode(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}

func deepCopy(src []int) []int {
    res := make([]int, len(src))
    for i := range src {
        res[i] = src[i]
    }

    return res
}
