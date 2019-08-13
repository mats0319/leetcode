package pn

func isPalindrome(x int) (isPalindrome bool) {
	// negative integer or integer end with '0' is not palindrome.
	if x < 0 || (x%10 == 0 && x != 0) {
		return
	}

	var reverse int

    for x > reverse {
        reverse = reverse*10 + x%10
        x /= 10
    }

    if x == reverse || x == reverse/10 {
        isPalindrome = true
    }

	return
}
