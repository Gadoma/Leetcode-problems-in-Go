package goleetcode0191

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0
	str := fmt.Sprintf("%b", num)
	for _, v := range str {
		if v == 49 {
			cnt++
		}
	}
	return cnt
}
