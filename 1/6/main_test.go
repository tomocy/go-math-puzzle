package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n         int
		maxLength int
		expected  int
	}{
		{n: 5, maxLength: 8, expected: 8},
		{n: 20, maxLength: 1000, expected: 26882},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n, maxLength: %d, %d", test.n, test.maxLength), func(t *testing.T) {
			actual := solve(test.n, test.maxLength)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
