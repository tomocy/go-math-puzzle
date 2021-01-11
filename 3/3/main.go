package main

import "fmt"

func main() {}

var tiles = []tile{
	{
		w: 1, h: 1,
	},
	{
		w: 2, h: 2,
	},
	{
		w: 4, h: 2,
	},
	{
		w: 4, h: 4,
	},
}

func solve(w, h int) int {
	return put(make(map[putCtx]int), make([]int, w), h)
}

func put(cache map[putCtx]int, heights []int, h int) int {
	ctx := putCtx{
		heights: fmt.Sprint(heights),
		h:       h,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	minI, minH := min(heights)
	if minH == h {
		cache[ctx] = 1
		return cache[ctx]
	}

	for _, tile := range tiles {
		fit := true
		for x := 0; x < tile.w; x++ {
			i := minI + x
			if i > len(heights)-1 {
				fit = false
				break
			}
			if heights[i]+tile.h > h {
				fit = false
				break
			}
			if heights[i] != heights[minI] {
				fit = false
				break
			}
		}
		if !fit {
			continue
		}

		copied := make([]int, len(heights))
		copy(copied, heights)
		for x := 0; x < tile.w; x++ {
			copied[minI+x] += tile.h
		}
		cache[ctx] += put(cache, copied, h)
	}

	return cache[ctx]
}

func min(vs []int) (int, int) {
	var (
		minI int
		minV int
	)

	for i, v := range vs {
		if i == 0 || v < minV {
			minI = i
			minV = v
		}
	}

	return minI, minV
}

type putCtx struct {
	heights string
	h       int
}

type tile struct {
	w, h int
}
