package goleetcode0217

func containsDuplicate(nums []int) bool {
	counts := map[int]int{}

	for _, number := range nums {
		counts[number]++

		if counts[number] > 1 {
			return true
		}
	}

	return false
}
