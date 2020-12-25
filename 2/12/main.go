package main

func main() {}

func solve(hole, per int) int {
	return solveWithCache(make(map[golfCtx]int), hole, per)
}

func solveWithCache(cache map[golfCtx]int, hole, per int) int {
	ctx := golfCtx{
		hole: hole,
		per:  per,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if hole <= 0 || per <= 0 {
		cache[ctx] = 0
		return 0
	} else if hole == 1 && 1 <= per && per <= 5 {
		cache[ctx] = 1
		return 1
	} else {
		for i := 1; i <= 5; i++ {
			cache[ctx] += solveWithCache(cache, hole-1, per-i)
		}
	}

	return cache[ctx]
}

type golfCtx struct {
	hole int
	per  int
}
