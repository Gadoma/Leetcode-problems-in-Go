package goleetcode0001

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, number := range nums {
		if j, ok := hashMap[number]; ok {
			return []int{j, i}
		} else {
			hashMap[target-number] = i
		}
	}
	return []int{0, 0}
}
