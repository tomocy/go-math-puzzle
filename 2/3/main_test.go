package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		r        int
		h        int
		expected int
	}{
		{
			r: 6, h: 4, expected: 3,
		},
		{
			r: 50, h: 35, expected: 1855967520,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("r,h:%d,%d", test.r, test.h), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.r, test.h)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
