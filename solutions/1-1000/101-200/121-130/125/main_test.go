package mario

import "testing"

var testCase = []struct {
	In     string
	Expect bool
}{
	// test cases here
	{"A man, a plan, a canal: Panama", true},
	{".,", true},
}

func TestIsPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isPalindrome(tcs[i].In) != tcs[i].Expect {
			t.Errorf("is palindrome test failed on case: %d\n", i)
		}
	}
}
