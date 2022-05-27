package mario

import "testing"

func TestRangeModule(t *testing.T) {
	ins := Constructor()

	res := make([]bool, 0)
	ins.AddRange(10, 20)
	ins.RemoveRange(14, 16)
	res = append(res, ins.QueryRange(10, 14))
	res = append(res, ins.QueryRange(13, 15))
	res = append(res, ins.QueryRange(16, 17))

	except := []bool{true, false, true}
	if len(res) != len(except) {
		t.Logf("unmatched res amount")
		t.Fail()
	}

	for i := range res {
		if res[i] != except[i] {
			t.Logf("test range module faild, index: %d\n\twant: %t\n\tget : %t\n", i, except[i], res[i])
			t.Fail()
		}
	}
}

func TestRangeModule_Case2(t *testing.T) {
	ins := Constructor()

	res := make([]bool, 0)
	ins.AddRange(10, 180)
	ins.AddRange(150, 200)
	ins.AddRange(250, 500)
	res = append(res, ins.QueryRange(50, 100))
	res = append(res, ins.QueryRange(180, 300))
	res = append(res, ins.QueryRange(600, 1000))
	ins.RemoveRange(50, 150)
	res = append(res, ins.QueryRange(50, 100))

	except := []bool{true, false, false, false}
	if len(res) != len(except) {
		t.Logf("unmatched res amount")
		t.Fail()
	}

	for i := range res {
		if res[i] != except[i] {
			t.Logf("test range module faild, index: %d\n\twant: %t\n\tget : %t\n", i, except[i], res[i])
			t.Fail()
		}
	}
}

func TestRangeModule_Case3(t *testing.T) {
	ins := Constructor()

	res := make([]bool, 0)
	res = append(res, ins.QueryRange(21, 34))
	res = append(res, ins.QueryRange(27, 87))
	ins.AddRange(44, 53)
	ins.AddRange(69, 89)
	res = append(res, ins.QueryRange(23, 26))
	res = append(res, ins.QueryRange(80, 84))
	res = append(res, ins.QueryRange(11, 12))
	ins.RemoveRange(86, 91)
	ins.AddRange(87, 94)
	ins.RemoveRange(34, 52)
	ins.AddRange(1, 59)
	ins.RemoveRange(62, 96)
	ins.RemoveRange(34, 83)
	res = append(res, ins.QueryRange(11, 59))
	res = append(res, ins.QueryRange(59, 79))
	res = append(res, ins.QueryRange(1, 13))
	res = append(res, ins.QueryRange(21, 23))
	ins.RemoveRange(9, 61)
	ins.AddRange(17, 21)
	ins.RemoveRange(4, 8)
	res = append(res, ins.QueryRange(19, 25))
	ins.AddRange(71, 98)
	ins.AddRange(23, 94)
	ins.RemoveRange(58, 95)
	res = append(res, ins.QueryRange(12, 78))
	ins.RemoveRange(46, 47)
	ins.RemoveRange(50, 70)
	ins.RemoveRange(84, 91)
	ins.AddRange(51, 63)
	ins.RemoveRange(26, 64)
	ins.AddRange(38, 87)
	res = append(res, ins.QueryRange(41, 84))
	res = append(res, ins.QueryRange(19, 21))
	res = append(res, ins.QueryRange(18, 56))
	res = append(res, ins.QueryRange(23, 39))
	res = append(res, ins.QueryRange(29, 95))
	ins.AddRange(79, 100)
	ins.RemoveRange(76, 82)
	ins.AddRange(37, 55)
	ins.AddRange(30, 97)
	ins.AddRange(1, 36)
	res = append(res, ins.QueryRange(18, 96))
	ins.RemoveRange(45, 86)
	ins.AddRange(74, 92)
	res = append(res, ins.QueryRange(27, 78))
	ins.AddRange(31, 35)
	res = append(res, ins.QueryRange(87, 91))
	ins.RemoveRange(37, 84)
	ins.RemoveRange(26, 57)
	ins.AddRange(65, 87)
	ins.AddRange(76, 91)
	res = append(res, ins.QueryRange(37, 77))
	res = append(res, ins.QueryRange(18, 66))
	ins.AddRange(22, 97)
	ins.AddRange(2, 91)
	ins.RemoveRange(82, 98)
	ins.RemoveRange(41, 46)
	ins.RemoveRange(6, 78)
	res = append(res, ins.QueryRange(44, 80))
	ins.RemoveRange(90, 94)
	ins.RemoveRange(37, 88)
	ins.AddRange(75, 90)
	res = append(res, ins.QueryRange(23, 37))
	ins.RemoveRange(18, 80)
	ins.AddRange(92, 93)
	ins.AddRange(3, 80)
	res = append(res, ins.QueryRange(68, 86))
	ins.RemoveRange(68, 92)
	res = append(res, ins.QueryRange(52, 91))
	ins.AddRange(43, 53)
	ins.AddRange(36, 37)
	ins.AddRange(60, 74)
	ins.AddRange(4, 9)
	ins.AddRange(44, 80)
	ins.RemoveRange(85, 93)
	ins.RemoveRange(56, 83)
	ins.AddRange(9, 26)
	res = append(res, ins.QueryRange(59, 64))
	res = append(res, ins.QueryRange(16, 66))
	ins.RemoveRange(29, 36)
	ins.RemoveRange(51, 96)
	ins.RemoveRange(56, 80)
	ins.AddRange(13, 87)
	res = append(res, ins.QueryRange(42, 72))
	ins.RemoveRange(6, 56)
	res = append(res, ins.QueryRange(24, 53))
	ins.AddRange(43, 71)
	ins.RemoveRange(36, 83)
	ins.RemoveRange(15, 45)
	res = append(res, ins.QueryRange(10, 48))

	except := []bool{false, false, false, true, false, false, false, true, true, false, false, true, true, false, false, false, true, false, true, false, false, false, false, true, false, false, false, true, false, false}
	if len(res) != len(except) {
		t.Logf("unmatched res amount")
		t.Fail()
	}

	for i := range res {
		if res[i] != except[i] {
			t.Logf("test range module faild, index: %d\n\twant: %t\n\tget : %t\n", i, except[i], res[i])
			t.Fail()
		}
	}
}
