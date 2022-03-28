package goleetcode0242

func isAnagram(s string, t string) bool {
	sr := map[byte]int{}
	tr := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sr[s[i]] += 1
		tr[t[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if sr[s[i]] != tr[s[i]] {
			return false
		}
	}

	return true
}
