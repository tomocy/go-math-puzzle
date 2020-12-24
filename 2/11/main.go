package main

func main() {}

func solve(w, h, n int) int {
	return combination(w, h, n)
}

func recursive(cache map[shareCtx]int, w, h, n int) int {
	ctx := shareCtx{
		w: w, h: h,
		n: n,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if n == 1 {
		cache[ctx] = 1
	} else {
		for i := 1; i < h; i++ {
			cache[ctx] += recursive(cache, w, i, n-1)
		}
		for i := 1; i < w; i++ {
			cache[ctx] += recursive(cache, i, h, n-1)
		}
	}

	return cache[ctx]
}

func combination(w, h, n int) int {
	var x int

	for i := 0; i < n; i++ {
		j := n - 1 - i
		x += nCr(w-1, i) * nCr(h-1, j) * nCr(n-1, i)
	}

	return x
}

func nCr(n, r int) int {
	if r == 0 || r == n {
		return 1
	}
	return nCr(n-1, r-1) + nCr(n-1, r)
}

type shareCtx struct {
	w, h int
	n    int
}
