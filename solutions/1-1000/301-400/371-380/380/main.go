package mario

import "math/rand"

type RandomizedSet struct {
	m    map[int]int // value - index
	data []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		data: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if !ok {
		this.m[val] = len(this.data)
		this.data = append(this.data, val)
	}

	return !ok
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.m[val]
	if ok {
		this.m[this.data[len(this.data)-1]] = index
		this.data[index], this.data[len(this.data)-1] = this.data[len(this.data)-1], this.data[index]

		this.data = this.data[:len(this.data)-1]
		delete(this.m, val)
	}

	return ok
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}
