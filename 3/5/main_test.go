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
			n: 3, expected: 2,
		},
		{
			n: 4, expected: 12,
		},
		{
			n: 5, expected: 312,
		},
		{
			n: 7, expected: 416880,
		},
		{
			n: 13, expected: 371511874881100800,
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
