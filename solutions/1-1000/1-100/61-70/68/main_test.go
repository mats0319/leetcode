package mario

import "testing"

type In struct {
	StringSlice0 []string
	Int0         int
}

var testCase = []struct {
	In     In
	Expect []string
}{
	// test cases here
	{In{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
		[]string{"This    is    an", "example  of text", "justification.  "}},
	{In{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
		[]string{"What   must   be", "acknowledgment  ", "shall be        "}},
	{In{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a",
		"computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
		[]string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
			"everything  else  we", "do                  "}},
}

func TestFullJustify(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !assertOnStringSlice(fullJustify(tcs[i].In.StringSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("full justify test failed on case: %d\n", i)
		}
	}
}

func assertOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
