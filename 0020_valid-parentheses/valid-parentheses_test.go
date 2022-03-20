package goleetcode0020

import "testing"

func TestValidParentheses(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  bool
	}{
		"single pair":                  {input: "()", want: true},
		"all pairs in a row":           {input: "(){}[]", want: true},
		"simple nested":                {input: "([{}])", want: true},
		"simple multiple nested":       {input: "([{}[]()])", want: true},
		"complex multiple nested":      {input: "(([{}]({[()]}[{([])}])))", want: true},
		"single opening":               {input: "(", want: false},
		"single closing":               {input: "}", want: false},
		"simple mismatch":              {input: "(}", want: false},
		"inverted mismatch":            {input: "}[", want: false},
		"simple two openings in a row": {input: "[[", want: false},
		"pair and just opening":        {input: "[](", want: false},
		"few pairs and a switch":       {input: "[]()}{", want: false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := validParentheses(testCase.input)
			if got != testCase.want {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
