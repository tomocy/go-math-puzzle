package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		hole     int
		per      int
		expected int
	}{
		{
			hole: 3, per: 12, expected: 10,
		},
		{
			hole: 18, per: 72, expected: 2546441085,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("hole,per:%v,%v", test.hole, test.per), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.hole, test.per)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
