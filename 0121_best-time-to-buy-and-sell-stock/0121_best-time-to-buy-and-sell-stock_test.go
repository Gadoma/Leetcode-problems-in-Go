package goleetcode0121

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  int
	}{
		"simple example":         {input: []int{7, 1, 5, 3, 6, 4}, want: 5},
		"decreasing order":       {input: []int{7, 6, 4, 3, 1}, want: 0},
		"only a single element":  {input: []int{1}, want: 0},
		"two identical elements": {input: []int{1, 1}, want: 0},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := maxProfit(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
