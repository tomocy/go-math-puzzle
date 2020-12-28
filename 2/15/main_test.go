package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		cards    int
		expected int
	}{
		{
			cards: 4, expected: 3,
		},
		{
			cards: 8, expected: 28630,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.cards), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.cards)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}

func BenchmarkSolveWithRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveWithRec(5)
	}
}

func BenchmarkSolveWithPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveWithPerm(5)
	}
}
