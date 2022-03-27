package goleetcode0007

import "math"

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}

	return result
}

func reverse(num int) int {
	//special case
	if num == 0 {
		return 0
	}

	digits := []int{}
	carry := num
	negative := 1
	result := 0
	add := 0

	// determine sign
	if num < 0 {
		carry *= -1
		negative = -1
	}

	// get digits
	for carry > 0 {
		digits = append(digits, carry%10)
		carry /= 10
	}

	// get number
	j := len(digits) - 1
	for i := 0; i < len(digits); i++ {
		add = intPow(10, j) * digits[i]

		if num > 0 && result+add > math.MaxInt32 {
			return 0
		}

		if num < 0 && (result+add)*negative < math.MinInt32 {
			return 0
		}

		result += add
		j--
	}

	// return number with sign
	return result * negative
}
