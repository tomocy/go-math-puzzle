package main

func main() {}

func solve(colors int) int {
	return pick(make(map[pickCtx]int), colors, 0, 0)
}

const (
	prevUsedTwice = iota
	prevUsedOnce
)

func pick(cache map[pickCtx]int, unused, once, prev int) int {
	ctx := pickCtx{
		unused: unused,
		once:   once,
		prev:   prev,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if unused == 0 && once == 0 && prev == prevUsedTwice {
		cache[ctx] = 1
		return cache[ctx]
	}

	if unused > 0 {
		cache[ctx] += pick(cache, unused-1, once+prev, prevUsedOnce)
	}
	if once > 0 {
		cache[ctx] += once * pick(cache, unused, once+prev-1, prevUsedTwice)
	}

	return cache[ctx]
}

type pickCtx struct {
	unused int
	once   int
	prev   int
}

func pick2(colors int) int {
	if colors == 1 {
		return 0
	}
	if colors == 2 {
		return 1
	}

	return (2*(colors-1)+1)*pick2(colors-1) + pick2(colors-2)
}
