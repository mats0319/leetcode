package mario

import "testing"

func TestNumArray(t *testing.T) {
	numArrayIns := Constructor([]int{1, 3, 5})

	res := make([]int, 0, 2)
	res = append(res, numArrayIns.SumRange(0, 2))
	numArrayIns.Update(1, 2)
	res = append(res, numArrayIns.SumRange(0, 2))

	expected := []int{9, 8}
	if len(res) != len(expected) {
		t.Errorf("unmatched res length, get: %d, want: %d\n", len(res), len(expected))
	}

	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("unexpected res, index: %d, get: %d, want: %d\n", i, res[i], expected[i])
		}
	}
}
