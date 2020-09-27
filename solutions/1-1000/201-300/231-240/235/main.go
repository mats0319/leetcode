package mario

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    small, big := getValue(p, q)
    return dfs(root, small, big)
}

func dfs(root *TreeNode, small, big int) (res *TreeNode) {
    if root == nil {
        return nil
    }

    if root.Val > big {
        res = dfs(root.Left, small, big)
    } else if root.Val < small {
        res = dfs(root.Right, small, big)
    } else {
        return root
    }

    return res
}

func getValue(a, b *TreeNode) (small, big int) {
    if a.Val < b.Val {
        small = a.Val
        big = b.Val
    } else {
        small = b.Val
        big = a.Val
    }

    return
}
