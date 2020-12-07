package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		n        int
		p        []int
		expected []int
	}{
		{n: 10, p: []int{250, 200, 150}, expected: []int{4, 3, 3}},
		{
			n: 289,
			p: []int{
				5381733, 1308265, 1279594, 2333899, 1023119, 1123891, 1914039, 2916976, 1974255, 1973115, 7266534, 6222666, 13515271, 9126214, 2304264, 1066328,
				1154008, 786740, 834930, 2098804, 2031903, 3700305, 7483128, 1815865, 1412916, 2610353, 8839469, 5534800, 1364316, 963579, 573441, 694352,
				1921525, 2843990, 1404729, 755733, 976263, 1385262, 728276, 5101556, 832832, 1377187, 1786170, 1166338, 1104069, 1648177, 1433566,
			},
			expected: []int{
				12, 3, 3, 5, 3, 3, 4, 7, 5, 5, 16, 14, 29, 20, 5, 3,
				3, 2, 2, 5, 5, 8, 16, 4, 3, 6, 19, 12, 3, 3, 2, 2,
				5, 6, 3, 2, 3, 3, 2, 11, 2, 3, 4, 3, 3, 4, 3,
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.p), func(t *testing.T) {
			actual := solve(test.n, test.p)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
