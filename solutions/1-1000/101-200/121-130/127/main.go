package mario

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if contains(wordList, endWord) < 0 {
		return 0
	}
	if index := contains(wordList, beginWord); index >= 0 {
		wordList = append(wordList[:index], wordList[index+1:]...)
	}

	depth := 1
	layer := make([]string, 0, len(wordList))
	layer = append(layer, beginWord)

	stop := false
	existFlag := make([]bool, len(wordList)) // attention: define on length
	for len(layer) > 0 && !stop {
		depth++
		nextLayer := make([]string, 0, len(wordList))
		for len(layer) > 0 && !stop {
			v := layer[0]
			layer = layer[1:]

			for j, w := range wordList {
				diffCount := 0
				for i := 0; i < len(v) && diffCount <= 1; i++ {
					if v[i] != w[i] {
						diffCount++
					}
				}

				if diffCount == 1 {
					if w == endWord {
						stop = true
						break
					}

					if !existFlag[j] {
						nextLayer = append(nextLayer, w)
						existFlag[j] = true
					}
				}
			}
		}

		layer = nextLayer
	}

	if !stop {
		depth = 0
	}

	return depth
}

func contains(list []string, str string) int {
	index := -1
	for i, l := range list {
		if l == str {
			index = i
			break
		}
	}

	return index
}
