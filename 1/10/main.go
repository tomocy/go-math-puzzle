package main

import (
	"math"
)

func main() {}

func solve(n int, p []int) []int {
	x, ok := adams(n, p)
	if !ok {
		return nil
	}

	seats := make([]int, len(p))
	for i := range seats {
		seats[i] = ceil(p[i], x)
	}
	return seats
}

func adams(n int, p []int) (int, bool) {
	min, max := 1, max(p)

	for min < max {
		var sum int

		x := (min + max) / 2
		for _, pop := range p {
			sum += ceil(pop, x)
		}

		if sum > n {
			min = x
			continue
		}

		if sum < n {
			max = x
			continue
		}

		return x, true
	}

	return 0, false
}

func max(vs []int) int {
	var max int

	for i, v := range vs {
		if i == 0 {
			max = v
			continue
		}

		if v > max {
			max = v
		}
	}

	return max
}

func ceil(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}
