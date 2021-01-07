package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		digits   int
		ticks    int
		expected int
	}{
		{
			digits: 3, ticks: 6, expected: 10,
		},
		{
			digits: 10, ticks: 50, expected: 167729959,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("digits,ticks:%v,%v", test.digits, test.ticks), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.digits, test.ticks)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
