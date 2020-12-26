package main

func main() {}

func solve(width, members, delta int) int {
	return cut(make(map[cutCtx]int), width, members, delta, 0, width)
}

func cut(cache map[cutCtx]int, width, members, delta, max, min int) int {
	ctx := cutCtx{
		width:   width,
		members: members,
		delta:   delta,
		max:     max,
		min:     min,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if max-min > delta {
		cache[ctx] = 0
		return cache[ctx]
	}

	if members == 0 {
		if width == 0 {
			cache[ctx] = 1
			return cache[ctx]
		} else {
			cache[ctx] = 0
			return cache[ctx]
		}
	}

	for i := 0; i <= width; i++ {
		cache[ctx] += cut(cache, width-i, members-1, delta, maxx(max, i), minn(min, i))
	}

	return cache[ctx]
}

func maxx(a, b int) int {
	max := a
	if b > max {
		max = b
	}
	return max
}

func minn(a, b int) int {
	min := a
	if b < min {
		min = b
	}
	return min
}

type cutCtx struct {
	width   int
	members int
	delta   int
	max     int
	min     int
}
