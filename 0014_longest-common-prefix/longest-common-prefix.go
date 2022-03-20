package goleetcode0014

import "strings"

func longestCommonPrefix(words []string) string {
	longestPrefix := ""
	prefixLength := 1
	firstWord := words[0]

	for range firstWord {
		prefix := ""
		for _, word := range words {
			prefix = firstWord[:prefixLength]

			if !strings.HasPrefix(word, prefix) {
				return longestPrefix
			}
		}
		longestPrefix = prefix
		prefixLength++
	}

	return longestPrefix
}
