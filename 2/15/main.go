package main

import (
	"sort"
)

func main() {}

func solve(cards int) int {
	return solveWithRec(cards)
}

func solveWithRec(cards int) int {
	return countUnsortedWithRec(make(map[countCtx]int), cards, 1)
}

func countUnsortedWithRec(cache map[countCtx]int, cards, pos int) int {
	ctx := countCtx{
		cards: cards,
		pos:   pos,
	}

	if x, ok := cache[ctx]; ok {
		return x
	}

	if cards == 0 {
		cache[ctx] = 0
		return cache[ctx]
	}
	if pos > cards {
		cache[ctx] = factorial(cards) - 1
		return cache[ctx]
	}

	cache[ctx] = pos * countUnsortedWithRec(cache, cards-1, pos)
	cache[ctx] += (cards - pos) * countUnsortedWithRec(cache, cards-1, pos+1)

	return cache[ctx]
}

type countCtx struct {
	cards int
	pos   int
}

func solveWithPerm(cards int) int {
	return countUnsortedWithPerm(cards)
}

func countUnsortedWithPerm(cards int) int {
	var x int

	permutated := permutateCards(cards)
	for _, vs := range permutated {
		move(vs)
		if !sort.IntsAreSorted(vs) {
			x++
		}
	}

	return x
}

func permutateCards(cards int) [][]int {
	vs := make([]int, cards)
	for i := range vs {
		vs[i] = i + 1
	}

	return permutate(vs)
}

func permutate(vs []int) [][]int {
	x := make([][]int, 0, factorial(len(vs)))
	permutatePart(&x, vs, len(vs))
	return x
}

func permutatePart(x *[][]int, vs []int, length int) {
	if length == 1 {
		copied := make([]int, len(vs))
		copy(copied, vs)
		*x = append(*x, copied)
		return
	}

	for i := 0; i < length; i++ {
		permutatePart(x, vs, length-1)
		j := length % 2 * i
		vs[j], vs[length-1] = vs[length-1], vs[j]
	}
}

func factorial(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
		x *= i
	}
	return x
}

func move(vs []int) {
	for i := range vs {
		j := vs[i] - 1
		vs[i], vs[j] = vs[j], vs[i]
	}
}
