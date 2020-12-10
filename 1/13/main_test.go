package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		length   int
		n        int
		expected int
	}{
		{length: 2, n: 1, expected: 1},
		{length: 3, n: 1, expected: 2},
		{length: 3, n: 2, expected: 2},
		{length: 4, n: 3, expected: 4},
		{length: 9, n: 5, expected: 28692},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("len, n: %d, %d", test.length, test.n), func(t *testing.T) {
			actual := solve(test.length, test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
