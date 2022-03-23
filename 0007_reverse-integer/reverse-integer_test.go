package goleetcode0007

import "testing"

func TestReverse(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  int
	}{
		"simple palindrome number":   {input: 121, want: 121},
		"simple number":              {input: 123, want: 321},
		"negative number":            {input: -123, want: -321},
		"negative palindrome number": {input: -121, want: -121},
		"ending with zero":           {input: 120, want: 21},
		"negative ending with zero":  {input: -120, want: -21},
		"single digit":               {input: 5, want: 5},
		"zero":                       {input: 0, want: 0},
		"value overflow":             {input: 1534236469, want: 0},
		"value negative overflow":    {input: -1534236469, want: 0},
		"character overflow":         {input: 1534236469123456, want: 0},
		"negative bigger overflow":   {input: -1534236469123456, want: 0},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := reverse(testCase.input)
			if got != testCase.want {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
