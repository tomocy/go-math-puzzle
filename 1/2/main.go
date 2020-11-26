package main

import "math"

func main() {}

func solve(n, begin, end int) int {
	points := int(math.Abs(float64(begin-end))) - 1
	revPoints := n - points - 1 - 1
	return 1<<points + 1<<revPoints - 1
}
