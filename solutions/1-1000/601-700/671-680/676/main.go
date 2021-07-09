package mario

// 1 <= word.length <= 100
type MagicDictionary struct {
	words [101][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		words: [101][]string{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := range dictionary {
		length := len(dictionary[i])

		list := this.words[length]
		list = append(list, dictionary[i])

		this.words[length] = list
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	list := this.words[len(searchWord)]

	isExist := false
	for i := range list {
		if calcDiffChars(searchWord, list[i]) == 1 {
			isExist = true
			break
		}
	}

	return isExist
}

// a.length == b.length
func calcDiffChars(a, b string) int {
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
			if diff > 1 {
				break
			}
		}
	}

	return diff
}
