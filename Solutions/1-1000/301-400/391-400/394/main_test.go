package mario

import "testing"

var testCase = []struct {
	In     string
	Expect string
}{
	{"3[a]2[bc]", "aaabcbc"},
	{"3[a2[c]]", "accaccacc"},
	{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	{"100[leetcode]", "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode"},
}

func TestDecodeString(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if decodeString(tcs[i].In) != tcs[i].Expect {
			t.Errorf("decode string test failed on case: %d\n", i)
		}
	}
}
