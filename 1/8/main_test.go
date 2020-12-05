package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		width, height int
		expected      int
	}{
		{width: 3, height: 2, expected: 4},
		{width: 10, height: 10, expected: 19},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("width, height: %d, %d", test.width, test.height), func(t *testing.T) {
			actual := solve(test.width, test.height)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
