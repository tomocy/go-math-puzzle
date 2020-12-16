package main

func main() {}

func solve(r, h int) int {
	return nCr(make(map[combi]int), h-1, r-h)
}

func nCr(cache map[combi]int, n, r int) int {
	c := combi{
		n: n,
		r: r,
	}

	if x, ok := cache[c]; ok {
		return x
	}

	if r == 0 || r == n {
		cache[c] = 1
	} else {
		cache[c] = nCr(cache, n-1, r-1) + nCr(cache, n-1, r)
	}

	return cache[c]
}

type combi struct {
	n, r int
}
