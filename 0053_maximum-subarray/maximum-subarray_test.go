package goleetcode0053

import "testing"

func TestMaxSubArray(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  int
	}{
		"few middle elements":           {input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		"the only element":              {input: []int{1}, want: 1},
		"all elements":                  {input: []int{5, 4, -1, 7, 8}, want: 23},
		"the only positive element":     {input: []int{-1, 1}, want: 1},
		"few middle elements with zero": {input: []int{-2, 1, -3, 4, -1, 0, 2, 1, -5, 4}, want: 6},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := maxSubArray(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
