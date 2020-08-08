package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var nodes = make([]*TreeNode, 0)

func recoverTree(root *TreeNode)  {
    loopTree(root)

    bigFlag := make([]int, 0)
    for i := 0; i < len(nodes) - 1; i++ {
        if nodes[i].Val > nodes[i+1].Val {
            bigFlag = append(bigFlag, i, i+1)
        }
    }

    if len(bigFlag) > 2 {
        bigFlag = []int{bigFlag[0], bigFlag[3]}
    }

    nodes[bigFlag[0]], nodes[bigFlag[1]] = nodes[bigFlag[1]], nodes[bigFlag[0]]

    return
}

func loopTree(root *TreeNode) {
    if root == nil {
        return
    }

    loopTree(root.Left)

    nodes = append(nodes, root)

    loopTree(root.Right)
}
