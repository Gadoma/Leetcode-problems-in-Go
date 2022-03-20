package goleetcode0009

import "testing"

func TestPalindromeNumber(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  bool
	}{
		"simple palindrome number":  {input: 121, want: true},
		"another palindrome number": {input: 3443, want: true},
		"positive number":           {input: 1211, want: false},
		"negative number":           {input: -121, want: false},
		"short number":              {input: 10, want: false},
		"single digit":              {input: 5, want: true},
		"zero":                      {input: 0, want: true},
		"larger palindrome number":  {input: 123123545321321, want: true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := palindromeNumber(testCase.input)
			if got != testCase.want {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
