package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected int
	}{
		{
			n: 8, expected: 12,
		},
		{
			n: 12, expected: 20,
		},
		{
			n: 16, expected: 32,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
