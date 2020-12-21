package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected int
	}{
		{
			n: 12, expected: 6,
		},
		{
			n: 100, expected: 30904,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
