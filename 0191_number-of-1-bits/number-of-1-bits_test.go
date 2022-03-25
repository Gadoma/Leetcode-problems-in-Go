package goleetcode0191

import "testing"

func TestHammingWeight(t *testing.T) {
	testCases := map[string]struct {
		input uint32
		want  int
	}{
		"few ones":   {input: 00000000000000000000000000001011, want: 3},
		"single one": {input: 00000000000000000000000010000000, want: 1},
		"more ones":  {input: 00000000000000000000011111111101, want: 10},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := hammingWeight(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
