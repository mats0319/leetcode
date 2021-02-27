package mario

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    m := make(map[*TreeNode]*TreeNode)
    m[root] = nil

    nodeStack := make([]*TreeNode, 0)
    nodeStack = append(nodeStack, root)
    for len(nodeStack) > 0 {
        node := nodeStack[0]
        nodeStack = nodeStack[1:]

        if node.Left != nil {
            m[node.Left] = node
            nodeStack = append(nodeStack, node.Left)
        }
        if node.Right != nil {
            m[node.Right] = node
            nodeStack = append(nodeStack, node.Right)
        }
    }

    ancestor := make(map[int]bool)
    for p != nil {
        ancestor[p.Val] = true
        p = m[p]
    }

    var res *TreeNode
    for q != nil {
        if ancestor[q.Val] {
            res = q
            break
        }

        q = m[q]
    }

    return res
}
