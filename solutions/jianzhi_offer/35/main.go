package mario

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
    m := make(map[*Node]*Node)

    pre := &Node{}
    q := pre

    for p := head; p != nil; p = p.Next {
        q.Next = &Node{Val: p.Val}
        q = q.Next
        m[p] = q
    }
    q = pre.Next
    for p := head; p != nil; p = p.Next {
        q.Random = m[p.Random]
        q = q.Next
    }

    return pre.Next
}
