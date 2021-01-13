package main

import (
	"fmt"
)

func main() {}

func solve(n int) int {
	return seats(make(map[seatsCtx]int), n, 1<<n-1) * factorial(n-1)
}

func seats(cache map[seatsCtx]int, n, women int) int {
	ctx := seatsCtx{
		n:     n,
		women: women,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if women == 0 {
		cache[ctx] = 1
		return cache[ctx]
	}

	ones := countOnes(fmt.Sprintf("%b", women))
	manLeft, manRight := ones-1, ones%n

	for woman := 0; woman < n; woman++ {
		next := 1 << woman
		if women&next == 0 {
			continue
		}
		if woman == manLeft || woman == manRight {
			continue
		}

		cache[ctx] += seats(cache, n, women^next)
	}

	return cache[ctx]
}

func countOnes(n string) int {
	var x int

	for _, r := range n {
		if r != '1' {
			continue
		}
		x++
	}

	return x
}

type seatsCtx struct {
	n     int
	women int
}

func factorial(n int) int {
	x := 1

	for i := 1; i <= n; i++ {
		x *= i
	}

	return x
}
