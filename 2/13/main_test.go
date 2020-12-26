package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		width    int
		members  int
		delta    int
		expected int
	}{
		{
			width: 5, members: 3, delta: 1, expected: 3,
		},
		{
			width: 40, members: 20, delta: 10, expected: 1169801856636575,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("width,member,delta:%v,%v,%v", test.width, test.members, test.delta), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.width, test.members, test.delta)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
