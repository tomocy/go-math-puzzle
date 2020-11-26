package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n          int
		begin, end int
		expected   int
	}{
		{
			n: 5, begin: 1, end: 5, expected: 8,
		},
		{
			n: 29, begin: 1, end: 17, expected: 36863,
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			actual := solve(test.n, test.begin, test.end)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
