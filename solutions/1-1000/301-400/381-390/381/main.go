package mario

import (
	"math/rand"
)

type RandomizedCollection struct {
	array   []int
	indexes map[int]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		array:   make([]int, 0),
		indexes: make(map[int]map[int]bool), // value : index : [bool]
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.array = append(this.array, val)

	if this.indexes[val] == nil {
		this.indexes[val] = make(map[int]bool)
	}
	this.indexes[val][len(this.array)-1] = true

	return len(this.indexes[val]) <= 1
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if len(this.indexes[val]) <= 0 {
		return false
	}

	i := -1
	for k := range this.indexes[val] {
		i = k
		break
	}

	delete(this.indexes[this.array[i]], i)

	this.indexes[this.array[len(this.array)-1]][i] = true
	delete(this.indexes[this.array[len(this.array)-1]], len(this.array)-1)

	this.array[i], this.array[len(this.array)-1] = this.array[len(this.array)-1], this.array[i]
	this.array = this.array[:len(this.array)-1]

	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
