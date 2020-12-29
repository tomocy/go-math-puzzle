package main

func main() {}

func solve(n int) int {
	var x int

	for i := 2; i < n-1; i++ {
		x += nCr(n, i) * (i - 1) * (n - i - 1)
	}

	return x
}

func nCr(n, r int) int {
	if r == 0 || r == n {
		return 1
	}
	return nCr(n-1, r-1) + nCr(n-1, r)
}
