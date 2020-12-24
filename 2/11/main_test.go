package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		w, h     int
		n        int
		expected int
	}{
		{
			w: 4, h: 3, n: 3, expected: 16,
		},
		{
			w: 20, h: 20, n: 10, expected: 16420955656,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("w,h,n:%v,%v,%v", test.w, test.h, test.n), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.w, test.h, test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive(make(map[shareCtx]int), 6, 6, 3)
	}
}

func BenchmarkCombination(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combination(6, 6, 3)
	}
}
