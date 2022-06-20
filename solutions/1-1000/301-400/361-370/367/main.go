package mario

func isPerfectSquare(num int) bool {
	i := 1
	for i*i < num {
		i++
	}

	return i*i == num
}
