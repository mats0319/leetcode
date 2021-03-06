package mario

import "strconv"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	panic("implement me")
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	panic("implement me")
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	panic("implement me")
}

func deserialize(s string) *NestedInteger {
	list := makeNestedInteger(s).GetList()
	res := list[0]
	for i := 1; i < len(list); i++ {
		res.Add(*list[i])
	}

	return res
}

func makeNestedInteger(s string) *NestedInteger {
	res := &NestedInteger{}
	if len(s) < 1 {
		return res
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			endIndex := matchedBrackets(s, i+1)

			res.Add(*makeNestedInteger(s[i+1 : endIndex]))

			i = endIndex
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			endIndex := parseInt(s, i)

			value, _ := strconv.Atoi(s[i : endIndex+1])
			newNode := &NestedInteger{}
			newNode.SetInteger(value)

			res.Add(*newNode)

			i = endIndex
		}
	}

	return res
}

// matchedBrackets returns end index of char ']'
func matchedBrackets(str string, start int) (endIndex int) {
	count := 1
	for i := start; i < len(str); i++ {
		if str[i] == '[' {
			count++
		} else if str[i] == ']' {
			count--
			if count <= 0 {
				endIndex = i
				break
			}
		}
	}

	return
}

// parseInt returns end index of last valid num
func parseInt(str string, start int) int {
	endIndex := start + 1
	for endIndex < len(str) && '0' <= str[endIndex] && str[endIndex] <= '9' {
		endIndex++
	}

	return endIndex - 1
}
