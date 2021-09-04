package mario

func frequencySort(s string) string {
	value := make(map[byte]int, len(s)) // char - frequency
	for i := range s {
		value[s[i]]++
	}

	frequency := make(map[int][]byte, len(value)) // frequency - char
	for k, v := range value {
		frequency[v] = append(frequency[v], k)
	}

	var res []byte
	for len(frequency) > 0 {
		maxFrequency := 0
		for k := range frequency {
			if k > maxFrequency {
				maxFrequency = k
			}
		}

		byteSlice := frequency[maxFrequency]
		for i := range byteSlice {
			char := byteSlice[i]
			for j := 0; j < value[byteSlice[i]]; j++ {
				res = append(res, char)
			}
		}

		delete(frequency, maxFrequency)
	}

	return string(res)
}
