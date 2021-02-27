package mario

import "sort"

type MedianFinder struct {
	data sort.IntSlice
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		data: sort.IntSlice{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
}

func (this *MedianFinder) FindMedian() float64 {
	this.data.Sort()
	length := this.data.Len()
	evenFlag := 0
	if length%2 == 0 {
		evenFlag = 1
	}

	return float64(this.data[length/2-evenFlag]+this.data[length/2]) / float64(2)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
