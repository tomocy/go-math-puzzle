package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		colors   int
		expected int
	}{
		{
			colors:   3,
			expected: 5,
		},
		{
			colors:   11,
			expected: 4939227215,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.colors), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.colors)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
