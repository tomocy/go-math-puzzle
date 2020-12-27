package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		cards    int
		limit    int
		expected int
	}{
		{
			cards: 6, limit: 2, expected: 4,
		},
		{
			cards: 32, limit: 10, expected: 607836582,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("cards,limits:%v,%v", test.cards, test.limit), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.cards, test.limit)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
