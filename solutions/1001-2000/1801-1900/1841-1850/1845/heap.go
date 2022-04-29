package mario

import "github.com/pkg/errors"

type binaryMinimumHeapOnArray struct {
    array []int
    size  int
}

// NewBinaryMinimumHeapOnArray value is 1~n
func NewBinaryMinimumHeapOnArray(n int) *binaryMinimumHeapOnArray {
    heapIns := &binaryMinimumHeapOnArray{
        array: make([]int, 1, n+1), // deprecated first index
    }

    for i := 1; i <= n; i++ {
        heapIns.Push(i)
    }

    return heapIns
}

func (h *binaryMinimumHeapOnArray) Push(v int) {
    h.array = append(h.array, v)

    h.size++

    h.shiftUp()
}

func (h *binaryMinimumHeapOnArray) Pop() (int, error) {
    if h.size < 1 {
        return 0, errors.New("empty heap")
    }

    popValue := h.array[1]

    h.array[1] = h.array[h.size]

    h.size--

    h.shiftDown()

    return popValue, nil
}

// shiftUp shift last element up to it position
func (h *binaryMinimumHeapOnArray) shiftUp() {
    index := h.size

    for index != 1 && h.array[index] < h.array[index/2] { // loop when 'index' is not root and 'index' < 'index'.parent
        h.array[index], h.array[index/2] = h.array[index/2], h.array[index]

        index /= 2
    }
}

// shiftDown shift root element down, in each step, we swap root and its smaller child
func (h *binaryMinimumHeapOnArray) shiftDown() {
    for index := 1; index*2 <= h.size; { // loop when 'index' is not leaf node
        left := h.array[index]
        if left > h.array[index*2] {
            left = h.array[index*2]
        }

        right := h.array[index]
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
