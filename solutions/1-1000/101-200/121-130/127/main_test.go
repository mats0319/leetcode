package mario

import "testing"

type In struct {
	String1      string
	String0      string
	StringSlice0 []string
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	//{In{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}}, 5},
	//{In{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}}, 0},
	//{In{"hot", "dog", []string{"hot", "dog"}}, 0},
	{In{"hot", "dog", []string{"hot", "dog", "cog", "pot", "dot"}}, 3},
}

func TestLadderLength(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if ladderLength(tcs[i].In.String1, tcs[i].In.String0, tcs[i].In.StringSlice0) != tcs[i].Expect {
			t.Errorf("ladder length test failed on case: %d\n", i)
		}
	}
}
