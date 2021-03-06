package mario

func isOneBitCharacter(bits []int) bool {
	i := 0
	for ; i < len(bits)-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}

	return i == len(bits)-1
}
