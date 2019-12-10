package mario

func isPalindrome(x int) bool {
	// negative integer or integer end with '0' is not palindrome.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var reverse int

	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	return x == reverse || x == reverse/10
}
