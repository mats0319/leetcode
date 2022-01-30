package mario

func nextGreatestLetter(letters []byte, target byte) byte {
	existLetters := make([]bool, 26)
	for i := range letters {
		index := int(letters[i] - 'a')
		if !existLetters[index] {
			existLetters[index] = true
		}
	}

	index := target - 'a'

	var char byte
	for char == 0 {
		index = (index + 1) % 26

		if existLetters[index] {
			char = index
			break
		}
	}

	return char+'a'
}
