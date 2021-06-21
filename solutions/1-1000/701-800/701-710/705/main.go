package mario

const base = 1009

type MyHashSet struct {
	data []*listNode
}

type listNode struct {
	Value int
	Next  *listNode
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{data: make([]*listNode, base)}
}

func (this *MyHashSet) Add(key int) {
    if this.Contains(key) {
        return
    }

    list := this.data[key % base]
    this.data[key%base] = &listNode{Value: key, Next: list}
}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}

	list := this.data[key % base]
	pre := &listNode{Next: list}
	preCopy := pre
	for pre != nil {
	    curr := pre.Next
	    if curr != nil && curr.Value == key {
	        pre.Next = curr.Next
	        break
        }

	    pre = pre.Next
    }

    this.data[key % base] = preCopy.Next
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	list := this.data[key%base]
	if list == nil {
		return false
	}

	isExist := false
	for p := list; p != nil; p = p.Next {
		if p.Value == key {
			isExist = true
			break
		}
	}

	return isExist
}
