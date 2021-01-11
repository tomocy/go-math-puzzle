package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		w        int
		h        int
		expected int
	}{
		{
			w: 4, h: 2, expected: 6,
		},
		{
			w: 10, h: 10, expected: 2657272845090,
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
