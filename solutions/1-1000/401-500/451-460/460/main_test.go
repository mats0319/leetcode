package mario

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	res := make([]int, 0)

	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	res = append(res, lfu.Get(1))
	lfu.Put(3, 3)
	res = append(res, lfu.Get(2))
	res = append(res, lfu.Get(3))
	lfu.Put(4, 4)
	res = append(res, lfu.Get(1))
	res = append(res, lfu.Get(3))
	res = append(res, lfu.Get(4))

	if !assertIntSlice(res, []int{1, -1, 3, -1, 3, 4}) {
		t.Errorf("460 test case 1 failed\n")
	}
}

func TestCase2(t *testing.T) {
	res := make([]int, 0)

	lfu := Constructor(0)
	lfu.Put(0, 0)
	res = append(res, lfu.Get(0))

	if !assertIntSlice(res, []int{-1}) {
		t.Errorf("460 test case 2 failed\n")
	}
}

func TestCase3(t *testing.T) {
	res := make([]int, 0)

	lfu := Constructor(2)
	lfu.Put(3, 1)
	lfu.Put(2, 1)
	lfu.Put(2, 2)
	lfu.Put(4, 4)
	res = append(res, lfu.Get(2))

	if !assertIntSlice(res, []int{2}) {
		t.Errorf("460 test case 3 failed\n")
	}
}

func TestCase4(t *testing.T) {
	res := make([]int, 0)

	// can be fold in editor
	{
		lfu := Constructor(10)
		lfu.Put(10, 13)
		lfu.Put(3, 17)
		lfu.Put(6, 11)
		lfu.Put(10, 5)
		lfu.Put(9, 10)
		res = append(res, lfu.Get(13))
		lfu.Put(2, 19)
		res = append(res, lfu.Get(2))
		res = append(res, lfu.Get(3))
		lfu.Put(5, 25)
		res = append(res, lfu.Get(8))
		lfu.Put(9, 22)
		lfu.Put(5, 5)
		lfu.Put(1, 30)
		res = append(res, lfu.Get(11))
		lfu.Put(9, 12)
		res = append(res, lfu.Get(7))
		res = append(res, lfu.Get(5))
		res = append(res, lfu.Get(8))
		res = append(res, lfu.Get(9))
		lfu.Put(4, 30)
		lfu.Put(9, 3)
		res = append(res, lfu.Get(9))
		res = append(res, lfu.Get(10))
		res = append(res, lfu.Get(10))
		lfu.Put(6, 14)
		lfu.Put(3, 1)
		res = append(res, lfu.Get(3))
		lfu.Put(10, 11)
		res = append(res, lfu.Get(8))
		lfu.Put(2, 14)
		res = append(res, lfu.Get(1))
		res = append(res, lfu.Get(5))
		res = append(res, lfu.Get(4))
		lfu.Put(11, 4)
		lfu.Put(12, 24)
		lfu.Put(5, 18)
		res = append(res, lfu.Get(13))
		lfu.Put(7, 23)
		res = append(res, lfu.Get(8))
		res = append(res, lfu.Get(12))
		lfu.Put(3, 27)
		lfu.Put(2, 12)
		res = append(res, lfu.Get(5))
		lfu.Put(2, 9)
		lfu.Put(13, 4)
		lfu.Put(8, 18)
		lfu.Put(1, 7)
		res = append(res, lfu.Get(6))
		lfu.Put(9, 29)
		lfu.Put(8, 21)
		res = append(res, lfu.Get(5))
		lfu.Put(6, 30)
		lfu.Put(1, 12)
		res = append(res, lfu.Get(10))
		lfu.Put(4, 15)
		lfu.Put(7, 22)
		lfu.Put(11, 26)
		lfu.Put(8, 17)
		lfu.Put(9, 29)
		res = append(res, lfu.Get(5))
		lfu.Put(3, 4)
		lfu.Put(11, 30)
		res = append(res, lfu.Get(12)) // wrong
		lfu.Put(4, 29)
		res = append(res, lfu.Get(3))
		res = append(res, lfu.Get(9))
		res = append(res, lfu.Get(6))
		lfu.Put(3, 4)
		res = append(res, lfu.Get(1))
		res = append(res, lfu.Get(10))
		lfu.Put(3, 29)
		lfu.Put(10, 28)
		lfu.Put(1, 20)
		lfu.Put(11, 13)
		res = append(res, lfu.Get(3))
		lfu.Put(3, 12)
		lfu.Put(3, 8)
		lfu.Put(10, 9)
		lfu.Put(3, 26)
		res = append(res, lfu.Get(8))
		res = append(res, lfu.Get(7))
		res = append(res, lfu.Get(5))
		lfu.Put(13, 17)
		lfu.Put(2, 27)
		lfu.Put(11, 15)
		res = append(res, lfu.Get(12))
		lfu.Put(9, 19)
		lfu.Put(2, 15)
		lfu.Put(3, 16)
		res = append(res, lfu.Get(1))
		lfu.Put(12, 17)
		lfu.Put(9, 1)
		lfu.Put(6, 19)
		res = append(res, lfu.Get(4))
		res = append(res, lfu.Get(5))
		res = append(res, lfu.Get(5))
		lfu.Put(8, 1)
		lfu.Put(11, 7)
		lfu.Put(5, 2)
		lfu.Put(9, 28)
		res = append(res, lfu.Get(1))
		lfu.Put(2, 2)
		lfu.Put(7, 4)
		lfu.Put(4, 22)
		lfu.Put(7, 24)
		lfu.Put(9, 26)
		lfu.Put(13, 28)
		lfu.Put(11, 26)
	}

	if !assertIntSlice(res, []int{-1, 19, 17, -1, -1, -1, 5, -1, 12, 3, 5, 5, 1, -1, 30, 5, 30, -1,
		-1, 24, 18, 14, 18, 11, 18, -1, 4, 29, 30, 12, 11, 29, 17, -1, 18, -1, 20, 29, 18, 18, 20}) {

		t.Errorf("460 test case 4 failed\n")
	}
}

func assertIntSlice(res, expected []int) bool {
	if len(res) != len(expected) {
		return false
	}

	isEqual := true
	for i := range res {
		if res[i] != expected[i] {
			isEqual = false
			break
		}
	}

	if !isEqual {
		fmt.Println("expected:", expected)
		fmt.Println("got     :", res)
	}

	return isEqual
}
