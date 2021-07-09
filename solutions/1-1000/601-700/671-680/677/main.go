package mario

import (
    "strings"
)

type MapSum struct {
	data []*pair
}

type pair struct {
	key   string
	value int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		data: make([]*pair, 0),
	}
}

func (this *MapSum) Insert(key string, val int) {
    isExist := false
    for i := range this.data {
        if this.data[i].key == key {
            isExist = true
            this.data[i].value = val
            break
        }
    }

    if isExist {
        return
    }

	this.data = append(this.data, &pair{
		key:   key,
		value: val,
	})
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for i := range this.data {
		if strings.HasPrefix(this.data[i].key, prefix) {
			sum += this.data[i].value
		}
	}

	return sum
}
