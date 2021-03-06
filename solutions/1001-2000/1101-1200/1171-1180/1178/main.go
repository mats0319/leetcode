package mario

import (
	"math/bits"
)

// puzzles.length = 7
// only contains lower-case letters
func findNumOfValidWords(words []string, puzzles []string) []int {
	m := make(map[int]int)
	for i := 0; i < len(words); i++ {
		var constitute int
		for j := 0; j < len(words[i]); j++ {
			constitute |= 1 << (words[i][j] - 'a')
		}

		if bits.OnesCount(uint(constitute)) <= 7 {
			m[constitute]++
		}
	}
	// loop all valid puzzle rule for each puzzle, add it's times
	res := make([]int, len(puzzles))
	for i := range puzzles {
		first := 1 << (puzzles[i][0] - 'a')

		for loop := 0; loop < 1<<6; loop++ {
			var constitute int
			for j := 0; j < 6; j++ {
				if (loop>>j)&1 == 1 {
					constitute |= 1 << (puzzles[i][j+1] - 'a')
				}
			}

			res[i] += m[constitute|first]
		}
	}

	return res
}
