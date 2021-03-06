package mario

func sortColors(nums []int) {
	count := [3]int{}
	for i := range nums {
		count[nums[i]]++
	}

	index := 0
	for i, c := range count {
		for j := 0; j < c; j++ {
			nums[index] = i
			index++
		}
	}

	return
}

func sortColors_ScanOnce(nums []int) {
	zero := 0
	two := len(nums) - 1
	for i := 0; i < len(nums) && i <= two; i++ {
		switch nums[i] {
		case 0:
			nums[zero], nums[i] = nums[i], nums[zero]
			zero++
		case 2:
			nums[two], nums[i] = nums[i], nums[two]
			two--
			if nums[i] != 1 {
				i--
			}
		}
	}

	return
}
