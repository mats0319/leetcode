package mario

import "strings"

// input is always valid
func isValidSerialization(preorder string) bool {
	chars := strings.Split(preorder, ",")
	count := 1
	i := 0
	for i < len(chars) && count > 0 {
		if chars[i] == "#" {
			count--
		} else {
			count++
		}

		i++
	}

	return count == 0 && i == len(chars)
}
