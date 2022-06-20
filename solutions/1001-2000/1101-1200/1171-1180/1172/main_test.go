package mario

import "testing"

func TestConstructor(t *testing.T) {
	ins := Constructor(2)

	res := make([]int, 0)

	ins.Push(1)
	ins.Push(2)
	ins.Push(3)
	ins.Push(4)
	ins.Push(5)
	res = append(res, ins.PopAtStack(0))
	ins.Push(20)
	ins.Push(21)
	res = append(res, ins.PopAtStack(0))
	res = append(res, ins.PopAtStack(2))
	res = append(res, ins.Pop())
	res = append(res, ins.Pop())
	res = append(res, ins.Pop())
	res = append(res, ins.Pop())
	res = append(res, ins.Pop())

	expected := []int{2, 20, 21, 5, 4, 3, 1, -1}

	if len(res) != len(expected) {
		t.Log("unmatched data length")
		t.Fail()
	}

	for i := range res {
		if res[i] != expected[i] {
			t.Logf("index: %d, get: %d, want: %d\n", i, res[i], expected[i])
			t.Fail()
		}
	}
}
