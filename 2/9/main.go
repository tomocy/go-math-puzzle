package main

func main() {}

func solve(n int) int {
	if n%2 == 0 {
		return 0
	}
	return tree(make(map[int]int), (n-1)/2)
}

func tree(cache map[int]int, n int) int {
	if x, ok := cache[n]; ok {
		return x
	}

	if n == 0 {
		cache[n] = 1
	} else {
		for i := 0; i < n; i++ {
			cache[n] += tree(cache, i) * tree(cache, n-i-1)
		}
	}

	return cache[n]
}
