package goleetcode0242

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  bool
	}{
		"valid anagram with repeated letters": {input: []string{"anagram", "nagaram"}, want: true},
		"same words":                          {input: []string{"abc", "abc"}, want: true},
		"not an anagram":                      {input: []string{"rat", "car"}, want: false},
		"different length words":              {input: []string{"abc", "abcd"}, want: false},
		"not an anagram duplicate letters":    {input: []string{"abcc", "abca"}, want: false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := isAnagram(testCase.input[0], testCase.input[1])
			if got != testCase.want {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
