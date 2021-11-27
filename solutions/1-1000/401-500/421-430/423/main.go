package mario

var number = [][]byte{
	[]byte("zero"),
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

var charSlice []int

// s 恰好符合要求
func originalDigits(s string) string {
	charSlice = make([]int, 26) // init for leetcode
	for i := range s {
		charSlice[int(s[i]-'a')]++
	}

	res0 := do('z', '0')
	res2 := do('w', '2')
	res4 := do('u', '4')
	res6 := do('x', '6')
	res8 := do('g', '8')
	res1 := do('o', '1')
	res3 := do('t', '3')
	res5 := do('f', '5')
	res7 := do('s', '7')
	res9 := do('e', '9')

	return res0 + res1 + res2 + res3 + res4 + res5 + res6 + res7 + res8 + res9
}

func do(matchChar byte, writeChar byte) string {
	times := charSlice[int(matchChar-'a')]

	res := writeChars(writeChar, times)
	subTimes(int(writeChar-'0'), times)

	return res
}

// 0 times is valid
func writeChars(char byte, times int) string {
	res := make([]byte, 0, times)
	for i := 0; i < times; i++ {
		res = append(res, char)
	}

	return string(res)
}

func subTimes(index int, times int) {
	for i := range number[index] {
		charSlice[int(number[index][i]-'a')] -= times
	}
}
