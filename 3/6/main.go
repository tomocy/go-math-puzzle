package main

import (
	"math"
)

func main() {}

func solve(n int) int {
	if n == 0 {
		return 1
	}

	var p int
	for pow(3, p+1) <= n {
		p++
	}

	return pow(2, p) + solve(min(n-pow(3, p), pow(3, p)-1))
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func min(vs ...int) int {
	var min int

	for i, v := range vs {
		if i == 0 || v < min {
			min = v
		}
	}

	return min
}
