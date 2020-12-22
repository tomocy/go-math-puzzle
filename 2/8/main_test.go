package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected int
	}{
		{
			n: 1, expected: 9,
		},
		{
			n: 12, expected: 2754844344633,
		},
		{
			n: 20, expected: 2684875320241747718,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}

func BenchmarkSolveWithCache_cache_used(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := new(cache)
		solveWithCache(cache, 20)
	}
}

func BenchmarkSolveWithCache_cache_ignored(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := &cache{
			doIgnore: true,
		}
		solveWithCache(cache, 20)
	}
}
