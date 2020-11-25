package main

func main() {}

func solve(n int) int {
	var x int

	for rocks := 0; rocks <= n; rocks++ {
		for scissors := 0; (rocks + scissors) <= n; scissors++ {
			papers := n - rocks - scissors
			maxs := max([]int{
				rocks, scissors, papers,
			})
			if len(maxs) != 1 {
				continue
			}

			x++
		}
	}

	return x
}

func max(xs []int) []int {
	var max int
	for i, x := range xs {
		if i == 0 {
			max = x
			continue
		}

		if max < x {
			max = x
		}
	}

	maxs := make([]int, 0, len(xs))
	for _, x := range xs {
		if x != max {
			continue
		}

		maxs = append(maxs, x)
	}

	return maxs
}
