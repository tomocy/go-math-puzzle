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
		{
			n: 1, expected: 7,
		},
		{
			n: 12, expected: 93,
		},
		{
			n: 15, expected: 1,
		},
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
