package main

func main() {}

func solve(n, a, b int) int {
	return game(make(map[gameCtx]int), n, a, b)
}

func game(cache map[gameCtx]int, n, a, b int) int {
	ctx := gameCtx{
		n: n,
		a: a,
		b: b,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if a == 0 || b == 0 {
		cache[ctx] = 1
	} else if n == 0 {
		cache[ctx] = 0
	} else {
		cache[ctx] += game(cache, n-1, a+1, b-1)
		cache[ctx] += game(cache, n-1, a, b)
		cache[ctx] += game(cache, n-1, a-1, b+1)
	}

	return cache[ctx]
}

type gameCtx struct {
	n    int
	a, b int
}
