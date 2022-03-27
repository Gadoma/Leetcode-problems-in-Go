package goleetcode0190

import "testing"

func TestReverseBits(t *testing.T) {
	testCases := map[string]struct {
		input uint32
		want  uint32
	}{
		"few ones":           {input: 43261596, want: 964176192},
		"most ones":          {input: 4294967293, want: 3221225471},
		"most ones reversed": {input: 3221225471, want: 4294967293},
		"just zero":          {input: 0, want: 0},
	}

	for name, testCases := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := reverseBits(testCases.input); got != testCases.want {
				t.Errorf("Expected value %v, got %v", testCases.want, got)
			}
		})
	}
}
