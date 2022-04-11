package mario

import "github.com/pkg/errors"

// binaryMinimumHeap binary minimum binaryMinimumHeap
type binaryMinimumHeap struct {
	array []int
	size  int
}

func NewBinaryMinimumHeap(data ...int) *binaryMinimumHeap {
	heapIns := &binaryMinimumHeap{
		array: make([]int, 1, len(data)+1), // deprecated first index
	}

	for i := range data {
		heapIns.Push(data[i])
	}

	return heapIns
}

func (h *binaryMinimumHeap) Pop() (int, error) {
	if h.size < 1 {
		return -1, errors.New("empty heap")
	}

	popValue := h.array[1]

	h.array[1] = h.array[h.size]

	h.size--

	h.shiftDown()

	return popValue, nil
}

func (h *binaryMinimumHeap) Push(v int) {
	h.array = append(h.array, v)

	h.size++

	h.shiftUp()
}

// shiftUp shift last element up to it position
func (h *binaryMinimumHeap) shiftUp() {
	index := h.size

	for index != 1 && h.array[index] < h.array[index/2] { // loop when 'index' is not root and 'index' < 'index'.parent
		h.array[index], h.array[index/2] = h.array[index/2], h.array[index]

		index /= 2
	}
}

// shiftDown shift root element down to it position
// in each step, we swap root with its smaller child
func (h *binaryMinimumHeap) shiftDown() {
	for index := 1; index*2 <= h.size; { // loop when 'index' is not leaf node
		left := h.array[index]
		right := h.array[index]

		if left > h.array[index*2] {
			left = h.array[index*2]
		}
		if index*2+1 <= h.size && right > h.array[index*2+1] {
			right = h.array[index*2+1]
		}

		if left == h.array[index] && right == h.array[index] { // 'index' <= left child and right child(or not exist)
			break
		}

		// swap 'index' with its smaller child
		if left <= right {
			h.array[index], h.array[index*2] = h.array[index*2], h.array[index]
			index *= 2
		} else {
			h.array[index], h.array[index*2+1] = h.array[index*2+1], h.array[index]
			index = index*2 + 1
		}
	}
}
