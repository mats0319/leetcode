package mario

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	trieTreeIns := newTrieTree()
	for i := range dictionary {
		trieTreeIns.insert(dictionary[i])
	}

	words := strings.Split(sentence, " ")
	for i := range words {
		trie := trieTreeIns.shortestTrie(words[i])
		if len(trie) > 0 {
			words[i] = trie
		}
	}

	return strings.Join(words, " ")
}

type trieTree struct {
	isEnd bool
	next  []*trieTree
}

func newTrieTree() *trieTree {
	return &trieTree{
		next: make([]*trieTree, 26),
	}
}

func (t *trieTree) insert(str string) {
	p := t
	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'

		if p.next[index] == nil {
			p.next[index] = newTrieTree()
		}

		p = p.next[index]
	}

	p.isEnd = true
}

func (t *trieTree) shortestTrie(str string) string {
	p := t
	i := 0
	for ; i < len(str) && p != nil && !p.isEnd; i++ {
		index := str[i] - 'a'

		p = p.next[index]
	}

	res := ""
	if p != nil && p.isEnd {
		res = str[:i]
	}

	return res
}
