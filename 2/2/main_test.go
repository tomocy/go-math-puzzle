package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected int
	}{
		{
			n: 3, expected: 21,
		},
		{
			n: 16, expected: 1149133,
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
