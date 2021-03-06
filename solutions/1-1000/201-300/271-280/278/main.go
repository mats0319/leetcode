package mario

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(_ int) bool {
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		index := left + (right-left)/2
		if isBadVersion(index) {
			right = index
		} else {
			left = index + 1
		}
	}

	return left
}
