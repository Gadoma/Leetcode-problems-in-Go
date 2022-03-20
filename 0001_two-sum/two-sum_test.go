package goleetcode0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := map[string]struct {
		input  []int
		target int
		want   []int
	}{
		"first two numbers":             {input: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		"some other numbers":            {input: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		"two of the same only numbers":  {input: []int{3, 3}, target: 6, want: []int{0, 1}},
		"two of the same numbers":       {input: []int{7, 3, 1, 3}, target: 6, want: []int{1, 3}},
		"only negative numbers":         {input: []int{-7, -3, -1, -3}, target: -6, want: []int{1, 3}},
		"negative and positive numbers": {input: []int{-7, -3, 1, 2}, target: -1, want: []int{1, 3}},
		"negative, zero and positive":   {input: []int{-7, -3, 0, 2}, target: 2, want: []int{2, 3}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := twoSum(testCase.input, testCase.target)
			if !reflect.DeepEqual(testCase.want, got) {
				t.Fatalf("Expected value %v, got %v instead!", testCase.want, got)
			}
		})
	}
}
