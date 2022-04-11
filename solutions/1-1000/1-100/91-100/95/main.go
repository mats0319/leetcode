package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
    data := make([]int, n)
    for i := range data {
        data[i] = i+1
    }

    return buildAllBST(data)
}

func buildAllBST(v []int) []*TreeNode {
    if len(v) < 1 {
        return nil
    } else if len(v) == 1 {
        return []*TreeNode{{Val: v[0]}}
    }

    res := make([]*TreeNode, 0)
    for i := range v {
        left := buildAllBST(v[:i])
        if len(left) < 1 {
            left = append(left, nil)
        }

        right := buildAllBST(v[i+1:])
        if len(right) < 1 {
            right = append(right, nil)
        }

        for _, l := range left {
            for _, r := range right {
                res = append(res, &TreeNode{
                    Val: v[i],
                    Left: l,
                    Right: r,
                })
            }
        }
    }

    return res
}
