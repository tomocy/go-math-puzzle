package main

import "math"

func main() {}

func solve(n int) int {
	if n <= 1 {
		return n
	}

	return 7*solve(n-2) + 3*pow(7, n-1)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
