package mario

import "testing"

var testCase = []struct {
	In     int
	Except bool
}{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isPalindrome(tcs[i].In) != tcs[i].Except {
			t.Errorf("palindrome number test failed on case: %d", i)
		}
	}
}
