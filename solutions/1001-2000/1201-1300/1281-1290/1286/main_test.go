package mario

import "testing"

func TestCase1(t *testing.T) {
	//ci := Constructor("abc", 2)
	//
	//next := make([]string, 0)
	//hasNext := make([]bool, 0)
	//
	//next = append(next, ci.Next())
	//hasNext = append(hasNext, ci.HasNext())
	//next = append(next, ci.Next())
	//hasNext = append(hasNext, ci.HasNext())
	//next = append(next, ci.Next())
	//hasNext = append(hasNext, ci.HasNext())
	//
	//if !assertStringSlice(next, []string{"ab", "ac", "bc"}) {
	//	t.Errorf("expect: [ab ac bc]\n")
	//	t.Errorf("got   : %v\n", next)
	//}
	//
	//if !assertBoolSlice(hasNext, []bool{true, true, false}) {
	//	t.Errorf("expect: [true true false]\n")
	//	t.Errorf("got   : %v\n", hasNext)
	//}
}

func assertStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}

func assertBoolSlice(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := range a {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
