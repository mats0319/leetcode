package mario

// some description: index given is from '0' below
type MyLinkedList struct {
    head *listNode
    length int // max valid index is length-1
}

type listNode struct {
    Val int
    Next *listNode
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{
        head:   nil,
        length: 0,
    }
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index > this.length-1 {
        return -1
    }

    p := this.head
    for count := 0; count < index; count++ { // p is always not nil
        p = p.Next
    }

    return p.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.head = &listNode{Val: val, Next: this.head}

    this.length++

    return
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    if this.length == 0 {
        this.AddAtHead(val)
        return
    }

    p := this.head
    for p.Next != nil {
        p = p.Next
    }

    p.Next = &listNode{Val: val}

    this.length++

    return
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.length {
        return
    } else if index <= 0 {
        this.AddAtHead(val)
        return
    } else if index == this.length {
        this.AddAtTail(val)
        return
    }

    pre := &listNode{Next: this.head}
    for count := 0; count < index; count++ {
        pre = pre.Next
    }

    pre.Next = &listNode{Val: val, Next: pre.Next}

    this.length++

    return
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.length {
        return
    }

    pre := &listNode{Next: this.head}
    preCopy := pre
    for count := 0; count < index; count++ {
        pre = pre.Next
    }

    pre.Next = pre.Next.Next
    this.head = preCopy.Next

    this.length--

    return
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
