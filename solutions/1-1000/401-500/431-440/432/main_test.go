package mario

import "testing"

func TestCase_1(t *testing.T) {
	res := make([]string, 0)

	ins := Constructor()
	ins.Inc("hello")
	ins.Inc("hello")
	res = append(res, ins.GetMaxKey())
	res = append(res, ins.GetMinKey())
	ins.Inc("leet")
	res = append(res, ins.GetMaxKey())
	res = append(res, ins.GetMinKey())

	resExpect := []string{"hello", "hello", "hello", "leet"}
	if len(res) != len(resExpect) {
		t.Errorf("> q 432, case 1 failed. %d - %d", len(res), len(resExpect))
	}

	for i := range res {
		if res[i] != resExpect[i] {
			t.Errorf("> q 432, case 1 failed, index: %d", i)
		}
	}
}
