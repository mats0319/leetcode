package mario

type Trie struct {
	isEnd bool
	next  []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next: make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if p.next[index] == nil {
			p.next[index] = &Trie{next: make([]*Trie, 26)}
		}

		p = p.next[index]
	}

	p.isEnd = true

	return
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isExist := true

	p := this
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if p.next[index] == nil {
			isExist = false
			break
		}

		p = p.next[index]
	}

	return isExist && p.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	isValid := true

	p := this
	for i := 0; i < len(prefix); i++ {
		index := int(prefix[i] - 'a')
		if p.next[index] == nil {
			isValid = false
			break
		}

		p = p.next[index]
	}

	return isValid
}
