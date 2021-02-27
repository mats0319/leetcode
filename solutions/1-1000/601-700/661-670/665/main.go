package mario

func checkPossibility(nums []int) bool {
    if len(nums) < 3 {
        return true
    }

	modifyChance := 1
	isValid := true
	if nums[0] > nums[1] {
	    nums[0] = nums[1]
	    modifyChance--
    }

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modifyChance <= 0 {
				isValid = false
				break
			}

			if nums[i-1] <= nums[i+1] {
			    nums[i] = nums[i+1]
            } else {
                nums[i+1] = nums[i]
            }

			modifyChance--

			i -= 2
		}
	}

	return isValid
}
