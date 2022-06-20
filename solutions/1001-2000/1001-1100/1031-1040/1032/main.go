package mario

type StreamChecker struct {
	dictionary *trie
	stack      []byte
}

// Constructor only contains small-case letters
func Constructor(words []string) StreamChecker {
	dictionary := &trie{
		children: make([]*trie, 27),
	}

	for i := range words {
		dictionary.insert(words[i])
	}

	return StreamChecker{
		dictionary: dictionary,
		stack:      make([]byte, 0),
	}
}

func (s *StreamChecker) Query(letter byte) bool {
	s.stack = append(s.stack, letter)

	containsFlag := false

	p := s.dictionary
	stackIndex := len(s.stack) - 1
	for p != nil && stackIndex >= 0 {
		index := s.stack[stackIndex] - 'a'

		p = p.children[index]
		if p == nil {
			break
		}

		if p.children[26] != nil { // matched
			containsFlag = true
			break
		}

		stackIndex--
	}

	return containsFlag
}

type trie struct {
	children []*trie // length: 27, if last item != nil, means it is a word
}

// insert in opposite order
func (t *trie) insert(v string) {
	p := t

	for i := len(v) - 1; i >= 0; i-- {
		index := v[i] - 'a'

		if p.children[index] == nil {
			p.children[index] = &trie{
				children: make([]*trie, 27),
			}
		}

		p = p.children[index]
	}

	p.children[26] = &trie{}
}
