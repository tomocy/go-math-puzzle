package main

func main() {}

func solve(n int) int {
	var x int

	for i := 1; i < n; i++ {
		x += i * patterns(n, i)
	}

	return x
}

func patterns(n, k int) int {
	return (n - k) * nPk(n, k-1)
}

func nPk(n, k int) int {
	x := 1

	for i := 0; i < k; i++ {
		x *= n - i
	}

	return x
}
