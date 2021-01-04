package main

func main() {}

func solve(n int) int {
	return move(make(map[moveCtx]int), n-2, 0)
}

func move(cache map[moveCtx]int, fwd, bwd int) int {
	ctx := moveCtx{
		fwd: fwd, bwd: bwd,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if fwd == 0 && bwd == 0 {
		cache[ctx] = 1
		return cache[ctx]
	}
	if fwd == 0 {
		cache[ctx] = 0
		return cache[ctx]
	}

	for i := 1; i <= fwd; i++ {
		cache[ctx] += move(cache, bwd+i-1, fwd-i)
	}

	return cache[ctx]
}

type moveCtx struct {
	fwd, bwd int
}
