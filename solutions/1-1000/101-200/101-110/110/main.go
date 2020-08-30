package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    _, isNotBalanced := deepth(root)

    return !isNotBalanced
}

func deepth(root *TreeNode) (deep int, isNotBalanced bool) {
    if root == nil {
        return 0, false
    }

    if isNotBalanced {
        return 0, true
    }

    leftDeep, isNotBalanced := deepth(root.Left)

    if isNotBalanced {
        return 0, true
    }

    rightDeep, isNotBalanced := deepth(root.Right)

    if isNotBalanced {
        return 0, true
    }

    max, sub := maxAndSub(leftDeep, rightDeep)

    return max, sub > 1
}

func maxAndSub(a, b int) (max int, sub int) {
    if a > b {
        max = a
        sub = a-b
    } else {
        max = b
        sub = b-a
    }

    return
}
