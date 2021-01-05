package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		members  int
		expected int
	}{
		{
			members: 32, expected: 1143455572,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.members), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.members)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
