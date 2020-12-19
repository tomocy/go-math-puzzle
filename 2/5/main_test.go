package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		pages    int
		days     int
		expected int
	}{
		{
			pages: 100, days: 1, expected: 1,
		},
		{
			pages: 100, days: 2, expected: 50,
		},
		{
			pages: 100, days: 3, expected: 834,
		},
		{
			pages: 180, days: 14, expected: 140615467,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("pages,days:%d,%d", test.pages, test.days), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.pages, test.days)
			if actual != test.expected {
				t.Errorf("\nexpected %d\n     got %d", test.expected, actual)
			}
		})
	}
}
