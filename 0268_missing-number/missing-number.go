package goleetcode0268

func missingNumber(nums []int) int {
	check := map[int]int{}
	result := 0

	for _, v := range nums {
		check[v] = 1
	}

	for i := 0; i < len(nums)+1; i++ {
		if check[i] == 0 {
			return i
		}
	}

	return result
}
