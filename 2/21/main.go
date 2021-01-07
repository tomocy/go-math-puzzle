package main

func main() {}

func solve(digits, ticks int) int {
	return tick(make(map[tickCtx]int), digits, ticks)
}

func tick(cache map[tickCtx]int, digits, ticks int) int {
	ctx := tickCtx{
		digits: digits,
		ticks:  ticks,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if digits < 0 || ticks < 0 {
		cache[ctx] = 0
		return cache[ctx]
	}
	if digits == 0 && ticks == 0 {
		cache[ctx] = 1
		return cache[ctx]
	}

	for i := 1; i <= 9; i++ {
		cache[ctx] += tick(cache, digits-1, ticks-i)
	}

	return cache[ctx]
}

type tickCtx struct {
	digits int
	ticks  int
}
