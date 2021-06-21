package mario

const base = 1009

type MyHashMap struct {
	data []*listNode
}

type listNode struct {
	Key   int
	Value int
	Next  *listNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{data: make([]*listNode, base)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
    list := this.data[key%base]

    if list == nil || !this.contains(key) {
        this.data[key%base] = &listNode{Key: key, Value: value, Next: list}
        return
    }

    for list != nil {
        if list.Key == key {
            list.Value = value
            break
        }

        list = list.Next
    }
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if !this.contains(key) {
		return -1
	}

	res := 0
	for p := this.data[key%base]; p != nil; p = p.Next {
		if p.Key == key {
			res = p.Value
			break
		}
	}

	return res
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	list := this.data[key%base]
	pre := &listNode{Next: list}
	preCopy := pre
	for pre != nil {
		curr := pre.Next
		if curr != nil && curr.Key == key {
			pre.Next = curr.Next
			break
		}

		pre = pre.Next
	}

	this.data[key%base] = preCopy.Next
}

func (this *MyHashMap) contains(key int) bool {
	list := this.data[key%base]
	if list == nil {
		return false
	}

	isExist := false
	for p := list; p != nil; p = p.Next {
		if p.Key == key {
			isExist = true
			break
		}
	}

	return isExist
}
