package sort

import (
	"math"
)

// BubbleSort bubble sort
func BubbleSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var (
		ordered bool
		border  int
		length  = len(is)
		pos     = length
	)

	for i := 0; i < length && !ordered; i++ {
		ordered = true
		border = small(pos, length-i-1) // border can't omit
		for j := 0; j < border; j++ {
			if is[j] > is[j+1] {
				ordered = false
				pos = j
				is[j], is[j+1] = is[j+1], is[j]
			}
		}
	}

	return
}

func small(a, b int) (small int) {
	if a < b {
		small = a
	} else {
		small = b
	}

	return
}

// SelectionSort selection sort
func SelectionSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var (
		ordered  bool
		length   = len(is)
		maxIndex int
	)

	for i := 0; i < length && !ordered; i++ {
		ordered = true
		maxIndex = 0
		for j := 0; j < length-i; j++ {
			if is[j] >= is[maxIndex] {
				ordered = false
				maxIndex = j
			}
		}
		if maxIndex != length-i-1 {
			is[maxIndex], is[length-i-1] = is[length-i-1], is[maxIndex]
		}
	}

	return
}

// MergeSort merge sort
func MergeSort(is []int) { // todo: think: if it can modify on 'is' directly?
	if len(is) <= 1 {
		return
	}

	const MAX = math.MaxInt64 // max type match elem in 'is'

	var (
		isCloned = append(is[:0:0], is...) // copy slice
		count    = 0
		length   = len(is)
	)
	{
		var (
			capacity = 1
		)

		for {
			count++
			capacity *= 2
			if capacity >= length {
				break
			}
		}

		for i := 0; i < capacity-length; i++ {
			isCloned = append(isCloned, MAX)
		}
	}

	scale := 1
	for i := 0; i < count; i++ {
		scale *= 2
		for j := 0; j < count-i; j++ { // group number = count - i
			BubbleSort(isCloned[j*scale : (j+1)*scale]) // bubble sort ^_^
		}
	}

	for i := 0; i < length; i++ {
		is[i] = isCloned[i]
	}

	return
}

// QuickSort quick sort
func QuickSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i] > is[0] {
				big++
			} else {
				small++
				if big != 0 {
					is[i], is[small] = is[small], is[i]
				}
			}
		}
		if small != 0 {
			is[0], is[small] = is[small], is[0]
		}
	}

	QuickSort(is[:small])
	QuickSort(is[small+1:])

	return
}

// CountingSort counting sort
func CountingSort(is []int) {
	var (
		countingArray []int
		min           int
	)
	{
		var max int
		{
			if len(is) <= 1 {
				return
			}

			max, min = is[0], is[0]
			for i := 1; i < len(is); i++ {
				if is[i] > max {
					max = is[i]
				} else if is[i] < min {
					min = is[i]
				}
			}
		}

		countingArray = make([]int, max-min+1)
		for i := 0; i < len(is); i++ {
			countingArray[is[i]-min]++
		}
	}

	var index int
	for i := 0; i < len(countingArray); i++ {
		for ; countingArray[i] > 0; countingArray[i]-- {
			is[index] = i + min
			index++
		}
	}

	return
}

// InsertionSort insertion sort
func InsertionSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var (
		pos, temp int
		length    = len(is)
	)
	for i := 1; i < length; i++ {
		if is[i] >= is[i-1] {
			continue
		}
		for pos = i; pos > 0 && is[i] < is[pos-1]; pos-- {
		}
		temp = is[i]
		for t := i; t > pos; t-- {
			is[t] = is[t-1]
		}
		is[pos] = temp
	}

	return
}
