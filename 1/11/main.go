package main

func main() {}

func solve(n int) int {
	if n <= 2 {
		return 1
	}

	x := 1
	y := solve(n - 1)
	x += y
	for i := 2; i < n; i++ {
		x += solve(i) * y
	}

	return x
}
