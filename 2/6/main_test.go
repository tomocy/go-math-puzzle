package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x, y     []int
		expected int
	}{
		{
			x: []int{3, 5, 0, 8, 1, 4, 2, 6, 7, 9}, y: []int{4, 8, 1, 7, 0, 6, 9, 2, 5, 3}, expected: 117,
		},
		{
			x: []int{8, 6, 8, 9, 3, 4, 1, 7, 6, 1}, y: []int{5, 1, 1, 9, 1, 6, 9, 0, 9, 6}, expected: 122,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("x,y:%v,%v", test.x, test.y), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.x, test.y)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
