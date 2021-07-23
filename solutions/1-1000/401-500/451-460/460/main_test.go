package mario

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	//res := make([]int, 0)
	//
	//lfu := Constructor(2)
	//lfu.Put(1, 1)
	//lfu.Put(2, 2)
	//res = append(res, lfu.Get(1))
	//lfu.Put(3, 3)
	//res = append(res, lfu.Get(2))
	//res = append(res, lfu.Get(3))
	//lfu.Put(4, 4)
	//res = append(res, lfu.Get(1))
	//res = append(res, lfu.Get(3))
	//res = append(res, lfu.Get(4))
	//
	//if !assertIntSlice(res, []int{1, -1, 3, -1, 3, 4}) {
	//	t.Errorf("460 test case 1 failed\n")
	//}
}

func TestCase2(t *testing.T) {
    //res := make([]int, 0)
	//
    //lfu := Constructor(0)
    //lfu.Put(0, 0)
    //res = append(res, lfu.Get(0))
	//
    //if !assertIntSlice(res, []int{-1}) {
    //    t.Errorf("460 test case 2 failed\n")
    //}
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
