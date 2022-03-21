package goleetcode0217

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := map[string]struct {
		input []int
		want  bool
	}{
		"few elements with duplicate": {input: []int{1, 2, 3, 1}, want: true},
		"few distinct elements":       {input: []int{1, 2, 3, 4}, want: false},
		"many elements":               {input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
		"two of the same":             {input: []int{1, 1}, want: true},
		"just a single element":       {input: []int{0}, want: false},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := containsDuplicate(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
