package goleetcode0007

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	//convert to string and extract sign
	runes := []rune(strconv.Itoa(x))
	prefix := ""

	if x < 0 {
		prefix = "-"
		runes = runes[1:]
	}

	// reverse order
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// check overflow
	if len(string(runes)) > 10 {
		return 0
	}
	value := prefix + fmt.Sprintf("%0*s", 10, string(runes))
	if value > strconv.Itoa(math.MaxInt32) {
		return 0
	}
	if x < 0 && value > strconv.Itoa(math.MinInt32) {
		return 0
	}

	// convert back to signed number
	ret, _ := strconv.Atoi(prefix + string(runes))

	return ret
}
