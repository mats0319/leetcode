package mario

import "testing"

func TestCase1(t *testing.T) {
	res := make([]bool, 0)

	mc := Constructor()
	res = append(res, mc.Book(20, 29))
	res = append(res, mc.Book(13, 22))
	res = append(res, mc.Book(44, 50))
	res = append(res, mc.Book(1, 7))
	res = append(res, mc.Book(2, 10))
	res = append(res, mc.Book(14, 20))
	res = append(res, mc.Book(19, 25))
	res = append(res, mc.Book(36, 42))
	res = append(res, mc.Book(45, 50))
	res = append(res, mc.Book(47, 50))
	res = append(res, mc.Book(39, 45))
	res = append(res, mc.Book(44, 50))
	res = append(res, mc.Book(16, 25))
	res = append(res, mc.Book(45, 50))
	res = append(res, mc.Book(45, 50))
	res = append(res, mc.Book(12, 20))
	res = append(res, mc.Book(21, 29))
	res = append(res, mc.Book(11, 20))
	res = append(res, mc.Book(12, 17))
	res = append(res, mc.Book(34, 40))
	res = append(res, mc.Book(10, 18))
	res = append(res, mc.Book(38, 44))
	res = append(res, mc.Book(23, 32))
	res = append(res, mc.Book(38, 44))
	res = append(res, mc.Book(15, 20))
	res = append(res, mc.Book(27, 33))
	res = append(res, mc.Book(34, 42))
	res = append(res, mc.Book(44, 50))
	res = append(res, mc.Book(35, 40))
	res = append(res, mc.Book(24, 31))

	if !assertBoolSlice(res, []bool{true, false, true, true, false, true, false, true, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}) {
		t.Errorf("729 test case 1 failed.")
	}
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
