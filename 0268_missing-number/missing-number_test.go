package goleetcode0268

import "testing"

func TestMissingNumber(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  int
	}{
		"small list":          {input: []int{3, 0, 1}, want: 2},
		"missing last number": {input: []int{0, 1}, want: 2},
		"bigger list":         {input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, want: 8},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := missingNumber(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
