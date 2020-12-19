package main

func main() {}

func solve(pages, days int) int {
	cache := make(map[plan]int)

	var x int

	for i := 1; i <= days; i++ {
		x += read(cache, pages, i, 0)
	}

	return x
}

func read(cache map[plan]int, pages, days, next int) int {
	p := plan{
		pages: pages,
		days:  days,
		next:  next,
	}

	if x, ok := cache[p]; ok {
		return x
	}

	if days == 1 {
		cache[p] = 1
	} else {
		min, max := next+1, (pages-days*(days-1)/2)/days
		for i := min; i <= max; i++ {
			cache[p] += read(cache, pages-i, days-1, i)
		}
	}

	return cache[p]
}

type plan struct {
	pages int
	days  int
	next  int
}
