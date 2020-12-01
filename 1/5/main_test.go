package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{n: 4, expected: 12},
		{n: 9, expected: 48},
		{n: 45, expected: 3518437540},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
