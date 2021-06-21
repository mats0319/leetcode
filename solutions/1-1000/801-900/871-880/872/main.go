package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return compareOnIntSlice(getLeaves(root1, make([]int, 0)), getLeaves(root2, make([]int, 0)))
}

func getLeaves(root *TreeNode, values []int) []int {
    if root == nil {
        return values
    } else if root.Left == nil && root.Right == nil {
        values = append(values, root.Val)
        return values
    }

    values = getLeaves(root.Left, values)

    values = getLeaves(root.Right, values)

    return values
}

// compareOnIntSlice return if array a and b is strict equal, include position and value
func compareOnIntSlice(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    isEqual := true
    for i := range a {
        if a[i] != b[i] {
            isEqual = false
            break
        }
    }

    return isEqual
}
