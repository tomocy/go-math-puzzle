package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected int
	}{
		{n: 10, k: 9, expected: 6},
		{n: 3000, k: 2500, expected: 897},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n,k:%d,%d", test.n, test.k), func(t *testing.T) {
			actual := solve(test.n, test.k)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
