package goleetcode0014

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  string
	}{
		"simple array with prefix":    {input: []string{"flower", "flow", "flight"}, want: "fl"},
		"simple array without prefix": {input: []string{"dog", "racecar", "car"}, want: ""},
		"simple mixed array":          {input: []string{"flower", "flow", "flight", "dog"}, want: ""},
		"array with empty string":     {input: []string{"flower", "flow", "flight", ""}, want: ""},
		"array of different letters":  {input: []string{"a", "b", "c"}, want: ""},
		"array of same letter":        {input: []string{"a", "a", "a"}, want: "a"},
		"array of empty strings":      {input: []string{"", "", ""}, want: ""},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := longestCommonPrefix(testCase.input)
			if got != testCase.want {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
