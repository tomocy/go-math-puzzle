package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		w, h     int
		expected int
	}{
		{
			w: 1, h: 1, expected: 4,
		},
		{
			w: 2, h: 2, expected: 8,
		},
		{
			w: 2, h: 3, expected: 14,
		},
		{
			w: 3, h: 3, expected: 20,
		},
		{
			w: 99, h: 101, expected: 20000,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("w,h:%v,%v", test.w, test.h), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.w, test.h)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
