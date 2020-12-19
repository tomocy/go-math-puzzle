package main

func main() {}

func solve(x, y []int) int {
	costs := make([][]int, len(y))
	for i := range y {
		costs[i] = make([]int, len(x))
		for j := range x {
			costs[i][j] = y[i] + x[j]

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				costs[i][j] += costs[i][j-1]
				continue
			}

			if j == 0 {
				costs[i][j] += costs[i-1][j]
				continue
			}

			costs[i][j] += min(costs[i-1][j], costs[i][j-1])
		}
	}

	return costs[len(y)-1][len(x)-1]
}

func min(a, b int) int {
	min := a
	if b < min {
		min = b
	}
	return min
}
