package main

func main() {}

func solve(w, h int) int {
	return change(make(map[changeCtx]int), w, h, w*h-1, 0)
}

func change(cache map[changeCtx]int, w, h, n, occupied int) int {
	ctx := changeCtx{
		w: w, h: h,
		n:        n,
		occupied: occupied,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if n < 0 {
		cache[ctx] = 1
		return cache[ctx]
	}

	row, col := n/w, n%w

	for i := 0; i < w*h; i++ {
		nextRow, nextCol := i/w, i%w
		if row == nextRow || col == nextCol {
			continue
		}
		if occupied&(1<<i) != 0 {
			continue
		}

		cache[ctx] += change(cache, w, h, n-1, occupied|1<<i)
	}

	return cache[ctx]
}

type changeCtx struct {
	w, h     int
	n        int
	occupied int
}
