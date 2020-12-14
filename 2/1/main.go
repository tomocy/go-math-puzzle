package main

func main() {}

func solve(n, cap int) int {
	return solveCache(make(map[int]int), n, cap)
}

func solveCache(cache map[int]int, n, cap int) int {
	if n == 0 {
		return 1
	}
	if cached, ok := cache[n]; ok {
		return cached
	}

	var x int

	for i := 1; i <= cap && i <= n; i++ {
		cache[n-i] = solveCache(cache, n-i, cap)
		x += cache[n-i]
	}

	return x
}
