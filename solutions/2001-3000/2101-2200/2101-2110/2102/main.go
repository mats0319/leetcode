package mario

type SORTracker struct {
	data  *scoreNode // pre
	times int        // get times
}

func Constructor() SORTracker {
	return SORTracker{
		data: &scoreNode{},
	}
}

func (s *SORTracker) Add(name string, score int) {
	spre := s.data
	sp := spre.next
	for sp != nil && score > sp.score {
		spre = spre.next
		sp = sp.next
	}

	if sp == nil || score < sp.score { // create new score node
		spre.next = &scoreNode{
			score:    score,
			nameList: &nameNode{},
			next:     sp,
		}
	}

	spre.next.length++

	npre := spre.next.nameList
	np := npre.next
	for np != nil && name > np.name {
		npre = npre.next
		np = np.next
	}

	npre.next = &nameNode{
		name: name,
		next: np,
	}
}

// Get always valid, ignore a lot valid-check in func
func (s *SORTracker) Get() string {
	s.times++
	index := s.times

	sp := s.data.next
	for sp != nil && sp.length < index {
		index -= sp.length
		sp = sp.next
	}

	np := sp.nameList
	for ; index > 0; index-- {
		np = np.next
	}

	return np.name
}

// scoreNode order by 'score' desc
type scoreNode struct {
	score int

	nameList *nameNode // pre
	length   int       // name list length

	next *scoreNode
}

// nameNode order by 'name' asc, use '>' on string type is ok
type nameNode struct {
	name string

	next *nameNode
}
