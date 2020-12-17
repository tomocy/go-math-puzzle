package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n          int
		start, end int
		expected   int
	}{
		{
			n: 4, start: 3, end: 2, expected: 3,
		},
		{
			n: 5, start: 2, end: 3, expected: 7,
		},
		{
			n: 5, start: 4, end: 2, expected: 4,
		},
		{
			n: 5, start: 2, end: 2, expected: 1,
		},
		{
			n: 15, start: 3, end: 10, expected: 1274766,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("n,start,end:%d,%d,%d", test.n, test.start, test.end), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n, test.start, test.end)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
