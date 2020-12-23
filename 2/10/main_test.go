package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		a, b     int
		expected int
	}{
		{
			n: 3, a: 1, b: 1, expected: 6,
		},
		{
			n: 24, a: 10, b: 10, expected: 1469180016,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("n,a,b:%d,%d,%d", test.n, test.a, test.b), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n, test.a, test.b)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
