package goleetcode0009

func palindromeNumber(x int) bool {
	//special cases
	switch {
	case x < 0:
		return false
	case x/10 < 1:
		return true
	case x%10 == 0:
		return false
	}

	digits := []int{}

	// split into digits
	for x >= 1 {
		digits = append(digits, x%10)
		x /= 10
	}

	// compare from beginning and end
	for i, j := 0, len(digits)-1; i < j; {
		if digits[i] != digits[j] {
			return false
		}

		i++
		j--
	}

	return true
}
