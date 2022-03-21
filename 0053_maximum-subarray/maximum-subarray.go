package goleetcode0053

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	max := nums[0]
	subMax := nums[0]

	for i := 1; i < len(nums); i++ {
		subMax = maxInt(subMax+nums[i], nums[i])
		max = maxInt(subMax, max)
	}

	return max
}
