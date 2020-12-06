package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		base      int
		maxLength int
		expected  []string
	}{
		{
			base: 3, maxLength: 3, expected: []string{"12", "22", "122"},
		},
		{
			base: 8, maxLength: 8, expected: []string{"24", "64", "134", "205", "463", "660", "661", "40663"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("base, maxLength: %d, %d", test.base, test.maxLength), func(t *testing.T) {
			actual := solve(test.base, test.maxLength)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
