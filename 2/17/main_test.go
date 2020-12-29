package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    int
		stock    int
		expected int
	}{
		{
			given: 5, stock: 10, expected: 3,
		},
		{
			given: 543210, stock: 987654, expected: 222223,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("given,stock:%v,%v", test.given, test.stock), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.given, test.stock)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
