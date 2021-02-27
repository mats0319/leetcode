package mario

type WordDictionary struct {
	data [][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		data: make([][]string, 501),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.data[len(word)] = append(this.data[len(word)], word)
}

func (this *WordDictionary) Search(word string) bool {
	if len(this.data[len(word)]) < 1 {
		return false
	}

	isExist := false

	for _, v := range this.data[len(word)] {
		isEqual := true
		for i := 0; i < len(word); i++ {
			if v[i] != word[i] && word[i] != '.' {
				isEqual = false
				break
			}
		}

		if isEqual {
			isExist = true
			break
		}
	}

	return isExist
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
