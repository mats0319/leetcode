package mario

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	res := make([]int, 0)

	data := Constructor(3, []int{4, 5, 8, 2})
	res = append(res, data.Add(3))
	res = append(res, data.Add(5))
	res = append(res, data.Add(10))
	res = append(res, data.Add(9))
	res = append(res, data.Add(4))

	fmt.Println("res     :", res)
	fmt.Println("expected: [4 5 5 8 8]")
}

func TestCase2(t *testing.T) {
	res := make([]int, 0)

	data := Constructor(1, []int{})
	res = append(res, data.Add(-3))
	res = append(res, data.Add(-2))
	res = append(res, data.Add(-4))
	res = append(res, data.Add(0))
	res = append(res, data.Add(4))

	fmt.Println("res     :", res)
	fmt.Println("expected: [-3 -2 -2 0 4]")
}

func TestCase3(t *testing.T) {
	res := make([]int, 0)

	data := Constructor(1, []int{-2})
	res = append(res, data.Add(-3))
	res = append(res, data.Add(0))
	res = append(res, data.Add(2))
	res = append(res, data.Add(-1))
	res = append(res, data.Add(4))

	fmt.Println("res     :", res)
	fmt.Println("expected: [-2 0 2 2 4]")
}
