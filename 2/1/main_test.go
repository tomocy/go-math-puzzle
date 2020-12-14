package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		cap      int
		expected int
	}{
		{
			n: 5, cap: 3, expected: 13,
		},
		{
			n: 32, cap: 6, expected: 1721441096,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("n,cap:%d,%d", test.n, test.cap), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n, test.cap)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
